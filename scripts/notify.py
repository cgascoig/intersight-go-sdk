from webexteamssdk import WebexTeamsAPI
import os

WEBEX_TEAMS_TOKEN = os.getenv("WEBEX_TEAMS_TOKEN")
DESTINATIONS = os.getenv("DESTINATIONS").split(",")


def send_notification(msg):
    api = WebexTeamsAPI(access_token=WEBEX_TEAMS_TOKEN)

    to_emails = DESTINATIONS

    for email in to_emails:
        api.messages.create(toPersonEmail=email, markdown=msg)

if __name__ == "__main__":
    print("Sending test message")
    send_notification("Hi, this is a **test** message.")