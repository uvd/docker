## Fluentd Cloudwatch

This image is used in our basic log alerting by grep implementation whilst maintaining output to Cloudwatch.

It extends the the official [kubernetes daemonset](https://github.com/fluent/fluentd-kubernetes-daemonset) image.

It adds the following plugins:
 * https://github.com/u-ichi/fluent-plugin-mail
 * https://github.com/sowawa/fluent-plugin-slack
 * https://github.com/sonots/fluent-plugin-grepcounter


Following on from the guide: https://docs.fluentd.org/v0.12/articles/splunk-like-grep-and-alert-email

Example configuration:

```
<source>
  @type tail
  path /var/log/containers/*.log
  pos_file /var/log/fluentd-containers.log.pos
  time_format %Y-%m-%dT%H:%M:%S.%NZ
  tag docker.*
  format json
  read_from_head true
</source>

<match docker.**>
  @type copy
  <store>
    @type cloudwatch_logs
    log_group_name "#{ENV['LOG_GROUP_NAME']}"
    auto_create_stream true
    use_tag_as_stream true
  </store>
  # Generic Exit|exit|EXIT
  <store>
    @type grepcounter
    count_interval 3 #Time window to grep and count the # of events
    input_key log
    regexp Exit|exit|EXIT
    threshold 1 #The # of events to trigger emitting an output
    add_tag_prefix alert
    delimiter \n
  </store>
</match>

<match alert.docker.**>
  @type copy
  <store>
    @type stdout
  </store>
  <store>
    @type slack
    webhook_url ****
    channel lm-alerts
    username fluentd
    icon_emoji :ghost:
    flush_interval 30s
    title %s
    title_keys input_tag
    message %s
    message_keys message
    mrkdwn true
  </store>
</match>
```