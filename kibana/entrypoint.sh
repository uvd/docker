#!/bin/bash

echo ""
echo "Installing Kibana plugins"
echo ""

IFS=';' read -r -a array <<< "${KIBANA_PLUGINS}"

for element in "${array[@]}"
do
    echo "Installing: $element"
    kibana-plugin install "$element"
done

echo ""
echo "Finished installing Kibana plugins"
echo ""

echo ""
echo "Starting Kibana"
echo ""
/usr/local/bin/kibana-docker
