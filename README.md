# Go automate for BrowserStack

https://automate.browserstack.com/

This is the automation for [BrowserStack](https://www.browserstack.com/).
Users can use this package to automate visiting any websites for any os and browsers for testing.

Below are five parameters for users to set in .env as their accounts of BrowserStack and Url that they would visit.

## Parameters
* <strong>BROWSERSTACK_USERNAME</strong>: BrowserStack user name
* <strong>BROWSERSTACK_ACCESS_KEY</strong>: BrowserStack access key
* <strong>BROWSERSTACK_URL</strong>: the url address users will visit
* <strong>BROWSERSTACK_HOSTS</strong>: binding of IP address and DNS
* <strong>BROWSERSTACK_RUN_LOCAL</strong>: bo0lean value, true if users want to visit interanl network for local teseting.
For the format, please refer to [.env.example](./env.example)

## Run Locally
If users want to visit internal address, e.g., http://localhost:8080, users need to start the local connection first.

* First visit here(https://www.browserstack.com/local-testing/app-automate) and go to <strong>command-line interface part</strong>
* Follow the instruction to install and execute the execution file by the following command
```
./BrowserStackLocal --key BROWSERSTACK_ACCESS_KEY
```
* Then users can execute the package to perform local testing.


## Run
```
make  
```
