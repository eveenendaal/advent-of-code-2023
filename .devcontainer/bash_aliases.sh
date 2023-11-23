#!/usr/bin/bash

# Aliases for gcloud
function gcplogin() {
    gcloud auth login --update-adc --no-launch-browser
}
