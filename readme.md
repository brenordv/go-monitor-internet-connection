# Internet Connection Monitor

## How it works
By default, this application will make a get request to ```https://clients3.google.com/generate_204``` every 30 seconds.
If the request don't go through, the app will assume the internet's out and start measuring how long the connection is out.
As soon as a request is made successfully, you'll receive a notification telling you how long the internet was out.

It's really simple: When your internet is down, this application will start a timer and when it's back up, it will 
send you a notification on Telegram, telling how long the internet was out. If you don't provide the required Telegram 
info, no notifications will be sent (but will always be logged).


## Changing URLs
This one is a special case. To use other urls, create a file called ```urls.txt``` with a single url each line and
place it in the same folder as the application.


## Configuration
You can overwrite default configurations using placing a file called ```runtime.config.json``` in the same folder 
as the application.


## Example file
```json
{
  "headers": {
    "Content-Type": "application/json"
  },
  "delayInSeconds": 30,
  "noConnInfoTTLinHours": 8760,
  "telegramChatId": 123123123123,
  "telegramBotToken": "token"  
}
```

- **headers**: Optional. Can pass an object to be used has custom headers. 
- **delayInSeconds**: Default: 30. This is the time between each request.
- **noConnInfoTTLinHours**: Everytime internet connection is restored, this application will save a log of your offline time. By default, they have a time to live of one year. To make the log last forever, change this value to -1.
- **telegramChatId**: if you want to get a notification when the internet connection is restored, inform the target chat_id in this property.
- **telegramBotToken**: if you want to get a notification when the internet connection is restored, inform the telegram bot token in this property. 


## How to use
### Local machine
Just run the application.
If you don't have an ```urls.txt``` file or ```runtime.config.json``` file, the application will run with the default values.
Note that, since it will not have the appropriate Telegram configuration, you will now receive any notifications when 
the internet connection is restored.

You can make it run as a service, so you don't have to manually execute this application everytime the machine starts.
Personally, i like to use it in a Docker container.

### Docker
I've included a ```Dockerfile``` and a bunch of scripts to build and run this application in a Docker container.


# Todo
- Implement fallback urls
- Create tests
- Commandline option to extract downtime logs