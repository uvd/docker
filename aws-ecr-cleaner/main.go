package main

import (
	"flag"
	"log"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// ECRImageSlice - Used to sort the AWS images response
type ECRImageSlice struct {
	Images []*ecr.ImageDetail
}

// NewECRImageSlice - Used to create a new sortable response of AWS ECR images
func NewECRImageSlice(images []*ecr.ImageDetail) *ECRImageSlice {
	return &ECRImageSlice{
		Images: images,
	}
}

func (slice ECRImageSlice) Len() int {
	return len(slice.Images)
}

func (slice ECRImageSlice) Less(i, j int) bool {
	t := *slice.Images[i].ImagePushedAt
	o := *slice.Images[j].ImagePushedAt
	return t.Before(o)
}

func (slice ECRImageSlice) Swap(i, j int) {
	slice.Images[i], slice.Images[j] = slice.Images[j], slice.Images[i]
}

func queryECR(svc *ecr.ECR, params *ecr.DescribeImagesInput) *ecr.DescribeImagesOutput {
	resp, err := svc.DescribeImages(params)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
		return nil
	}

	return resp
}

func main() {

	repository := flag.String("repository", "", "The repository to clean up.")
	numToKeep := flag.Int("keep", 10, "The amount of images to keep.")
	dryRun := flag.Bool("dry-run", false, "Don't actually remove the images. Used to see which ones will be deleted.")
	flag.Parse()

	sess := session.Must(session.NewSession())
	svc := ecr.New(sess)

	// Store slice of Image IDs to delete (use the digest)
	var imagesToDel []*ecr.ImageIdentifier

	// Find all dangling/untagged images
	untaggedParams := &ecr.DescribeImagesInput{
		RepositoryName: aws.String(*repository), // Required
		Filter: &ecr.DescribeImagesFilter{
			TagStatus: aws.String("UNTAGGED"),
		},
	}
	untaggedResp := queryECR(svc, untaggedParams)
	untaggedImages := untaggedResp.ImageDetails

	// Add all untagged images to be deleted
	for _, image := range untaggedImages {
		imagesToDel = append(imagesToDel, &ecr.ImageIdentifier{
			ImageDigest: image.ImageDigest,
		})
	}

	// Find all images
	allParams := &ecr.DescribeImagesInput{
		RepositoryName: aws.String(*repository),
	}
	allResp := queryECR(svc, allParams)
	allImages := NewECRImageSlice(allResp.ImageDetails)
	sort.Sort(sort.Reverse(allImages))

	// Add all but n newest images to be deleted
	if *numToKeep < len(allImages.Images) {
		for _, image := range allImages.Images[*numToKeep:] {
			imagesToDel = append(imagesToDel, &ecr.ImageIdentifier{
				ImageDigest: image.ImageDigest,
			})
		}
	}

	log.Printf("Repository: %s", *repository)
	log.Printf("Total images: %d", len(allImages.Images))
	log.Printf("Untagged images: %d", len(untaggedImages))
	log.Printf("Keeping %d images:", *numToKeep)
	if *numToKeep < len(allImages.Images) {
		log.Println(allImages.Images[:*numToKeep])
	} else {
		log.Println(allImages.Images)
	}
	log.Printf("Deleting %d images: ", len(imagesToDel))
	log.Println(allImages.Images[*numToKeep:])

	if !*dryRun && len(imagesToDel) > 0 {
		// Commit deleted images
		delParams := &ecr.BatchDeleteImageInput{
			ImageIds:       imagesToDel,
			RepositoryName: aws.String(*repository), // Required
		}
		delResp, err := svc.BatchDeleteImage(delParams)

		if err != nil {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Println(err.Error())
			return
		}

		// Pretty-print the response data.
		log.Println(delResp)
	}

}
