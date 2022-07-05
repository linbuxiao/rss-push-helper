# Feedly Helper

A simple feedly crawler.

## Usage

1. First you need to go [here](https://feedly.com/v3/auth/dev) to request a token, which will give you the ability to call the api.
2. Download the binary from the release that matches your operating system.
3. Use the following to make the call
    ```bash
   ./rss-push-helper --token <YOUR TOKEN>
    ```
4. You will see the updated feed titles and addresses from the current to the previous day range. 
By default it is rendered as html for sending to email or other. 
It can also be rendered to json using `-f json`.

