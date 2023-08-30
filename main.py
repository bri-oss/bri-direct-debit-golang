import json
import os
import requests
import datetime
import hashlib
import hmac

# Replace with your actual token and secret
token = os.environ.get("BRI_CLIENT_ID", "client_id")
secret = os.environ.get("BRI_CLIENT_SECRET", "client_secret")

account = "888801000157508"
url = "https://bri-partner-sandbox.deno.dev/v2/inquiry/" + account
headers = {
    "Authorization": f"Bearer {token}",
    "BRI-Signature": "",
    "BRI-Timestamp": "",
}

# Generate the timestamp in the required format
timestamp = datetime.datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%S.%f')[:-3] + "Z"
headers["BRI-Timestamp"] = timestamp

# Calculate the signature
string_to_sign = f"GET\n{timestamp}\n/v2/inquiry/888801000157508\n"
signature = hmac.new(secret.encode(), string_to_sign.encode(), hashlib.sha256).hexdigest()
headers["BRI-Signature"] = signature

response = requests.get(url, headers=headers)
response_json = response.json()

print("Credentials from Envars")
print("========================")
print("Client ID:", token)
print("CLient Secret:", secret, "\n")

print("Response from API")
print("=================")
print("Response status code:", response.status_code)
print("Response content:")
print(json.dumps(response_json, indent=4))
