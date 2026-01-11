import json
import urllib.request
import os
import sys

def get_latest_kernel():
    api_url = "https://www.kernel.org/releases.json"
    print(f"Fetching {api_url}...")

    try:
        with urllib.request.urlopen(api_url) as response:
            data = json.loads(response.read().decode())
    except Exception as e:
        print(f"Error fetching releases.json: {e}")
        sys.exit(1)

    latest_version = data.get("latest_stable", {}).get("version")
    if not latest_version:
        print("Could not find latest_stable version in JSON")
        sys.exit(1)

    print(f"Latest stable version: {latest_version}")

    download_url = None
    for release in data.get("releases", []):
        if release.get("version") == latest_version:
            download_url = release.get("source")
            break

    if not download_url:
        print(f"Could not find source URL for version {latest_version}")
        sys.exit(1)

    print(f"Download URL: {download_url}")

    filename = os.path.basename(download_url)
    if os.path.exists(filename):
        print(f"File {filename} already exists. Skipping download.")
        return

    print(f"Downloading {filename}...")
    try:
        # Download with progress
        def reporthook(blocknum, blocksize, totalsize):
            readsofar = blocknum * blocksize
            if totalsize > 0:
                percent = readsofar * 1e2 / totalsize
                s = "\r%5.1f%% %*d / %d" % (
                    percent, len(str(totalsize)), readsofar, totalsize)
                sys.stderr.write(s)
                if readsofar >= totalsize: # near the end
                    sys.stderr.write("\n")
            else: # total size is unknown
                sys.stderr.write("read %d\n" % (readsofar,))

        urllib.request.urlretrieve(download_url, filename, reporthook)
        print("Download complete!")
    except Exception as e:
        print(f"Error downloading file: {e}")
        sys.exit(1)

if __name__ == "__main__":
    get_latest_kernel()
