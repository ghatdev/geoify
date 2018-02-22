[![Build Status](https://travis-ci.org/ghatdev/geoify.svg?branch=master)](https://travis-ci.org/ghatdev/geoify)
# geoify
Simple Geo-Locate IP API service  


[![Donation](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=RX3YMBJTX35FC&lc=KR&item_name=geoify&currency_code=USD&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted)


This product includes GeoLite2 data created by MaxMind, available from [http://www.maxmind.com](http://www.maxmind.com)
[![GeoLite2 License](https://i.creativecommons.org/l/by-sa/4.0/88x31.png)](http://creativecommons.org/licenses/by-sa/4.0/)

## Usage
GET https://geoify.herokuapp.com/YOUR-IP-ADDRESS-TO-QUERY  
returns (locale: en)  
- City Name
- Subdivision Name
- Country Name
- ISO Country Code
- Time Zone
    
## Example
```bash
curl http://geoify.herokuapp.com/172.217.26.46
```
returns
```json
{"cityName":"Mountain View","subdivisionName":"California","countryName":"United States","isoCountryCode":"US","timeZone":"America/Los_Angeles"}
```
