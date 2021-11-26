#!/usr/bin/env python3

import requests
from bs4 import BeautifulSoup
import logging
import re
import os
import subprocess

import notify

def get_server_version():
    url = 'https://intersight.com/apidocs/downloads/'

    response = requests.get(url)
    if not response.ok:
        logging.error("Request failed")
        exit(1)

    soup = BeautifulSoup(response.text, 'html.parser')

    data_divs = soup.find_all(id="baseUrl")
    if len(data_divs) != 1:
        logging.error(f"Number of matching divs not exactly 1: {len(data_divs)}")
        exit(1)

    data_url = data_divs[0]['data-value']

    logging.info(f"Base data URL: {data_url}")

    m = re.match(r'https://cdn.intersight.com/components/an-apidocs/([0-9.\-]+)/', data_url)
    if m is None:
        logging.error(f"Unable to parse version from base data url: {data_url}")
        exit(1)
    
    return m.group(1)

def get_local_version():
    try:
        with open("OPENAPI_VERSION", "r") as f:
            version = f.readline().strip()
            return version
    except:   
        return "0"

def set_local_version(version):
    with open("OPENAPI_VERSION", "w") as f:
        f.write(version)

def download_spec(version, path):

    download_url = f"https://cdn.intersight.com/components/an-apidocs/{version}/model/intersight-openapi-v3-{version.replace('-','.')}.json"

    logging.info(f"Using download URL: {download_url}")

    response = requests.get(download_url)
    if not response.ok:
        logging.error("Download failed")
        exit(1)

    spec = re.sub(b'This document was created on \d\d\d\d-\d\d-\d\dT\d\d:\d\d:\d\d[A-Z].', b'', response.content)

    with open(path, 'wb') as f:
        f.write(spec)

def main():
    logging.getLogger().setLevel(logging.DEBUG)

    os.chdir("..")
    logging.info(f"Working directory: {os.getcwd()}")

    try:
        logging.info("Running git checkout main")
        subprocess.run(["git", "checkout", "main"], check=True)
        logging.info("Running git pull")
        subprocess.run(["git", "pull"], check=True)
    except:
        logging.error("Error occurred during git setup operations")
        notify.send_notification("Update intersight-go-sdk - error during git setup operations")
        exit(1)

    local_version = get_local_version()
    logging.info(f"Current local version: {local_version}")

    version = get_server_version()
    logging.info(f"Version found: {version}")

    if os.getenv("TEST") is not None:
        rand = random.randint(0,999999)
        version=f"{version}-test{rand}"
        branch=f"test{rand}"
        logging.info(f"Overridding version to {version} to force test update and using branch {branch}")
        res = subprocess.run(["git", "checkout", "-b", branch], check=True)

    if local_version == version:
        logging.info(f"No new version, exiting")
        notify.send_notification("Update intersight-go-sdk - no new version found")
        exit(0)

    logging.info("New version found")
    notify.send_notification(f"Update intersight-go-sdk - new version found ({version}")

    download_spec(version, "spec/intersight-openapi-v3.json")
    logging.info("Download successful")

    logging.info("Running 'make all'")
    res = subprocess.run(["make", "all"], env={**os.environ, "OPENAPI_VERSION": version})
    if res.returncode != 0:
        logging.error("Error occurred running make, reverting spec file and package")
        notify.send_notification("Update intersight-go-sdk - error during make")
        subprocess.run(["git", "checkout", "--", "spec/", "intersight/"])
        exit(1)
    
    logging.info("make ran successfully")
    notify.send_notification("Update intersight-go-sdk - make ran successfully")

    set_local_version(version)

    try:
        logging.info("Performing git operations")

        logging.info("Running git add ...")
        res = subprocess.run(["git", "add", "spec/", "intersight/", "OPENAPI_VERSION"], check=True)
        logging.info("Running git commit ...")
        res = subprocess.run(["git", "commit", "-m", f"Update to version {version}"], check=True)
        logging.info("Running git tag ...")
        res = subprocess.run(["git", "tag", f"intersight/v{version}"], check=True)
        logging.info("Running git push ...")
        res = subprocess.run(["git", "push", "origin"], check=True)
        logging.info("Running git push --tags...")
        res = subprocess.run(["git", "push", "origin", "--tags"], check=True)
    except:
        logging.error("Error occurred during git operations")
        notify.send_notification("Update intersight-go-sdk - error during git operations")
        exit(1)

    logging.info("Git operations completed successfully")

    notify.send_notification(f"Update intersight-go-sdk - finished updating to version {version}")
    logging.info("Finished")
    

if __name__ == "__main__":
    main()
