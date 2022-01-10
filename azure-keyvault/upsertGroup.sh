#!/bin/bash
groupName=$1
groupRegion=$2
echo "$groupName"
group=$(az group list --query "[?name == '$groupName'].name" -o tsv)
if [ "$group" != "$groupName" ]; then
    az group create -n $groupName -l $groupRegion
fi