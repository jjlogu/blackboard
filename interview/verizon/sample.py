import requests as r

r=redfish.get("storge")
retry = -1 # Null

# global network - 
# virtualization 5G 4G
try:
    resp = r.get("https://xyz.com/redfish/Managers", headers={})
    if resp.status_code != 200:
        pass
except:
    retry  += 1



