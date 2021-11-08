from webexteamssdk import WebexTeamsAPI
import creds

def send_notification(msg):
    api = WebexTeamsAPI(access_token=creds.WEBEX_TEAMS_TOKEN)

    to_emails = creds.DESTINATIONS

    for email in to_emails:
        api.messages.create(toPersonEmail=email, markdown=msg)

if __name__ == "__main__":
    print("Sending test message")
    send_notification("Hi, this is a **test** message.")