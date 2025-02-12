# Fitbit Workout Tracker sync

Command line utility to sync your Fitbit activities with your Workout Tracker
account.

:warning: **This is still a work in progress.**

## Summary

This is a separate tool which will access your Fitbit activities and upload them
(sync) to your Workout Tracker.

You need to register this application with Fitbit. The flow uses OAuth2 with
refresh tokens. The application takes care of refreshing this token.

Start by reading these resources:

- [Web API](https://dev.fitbit.com/build/reference/web-api/)
- [Developer Guide](https://dev.fitbit.com/build/reference/web-api/developer-guide/)

## Linking your Fitbit account

To link your Fitbit account to this application, this application will start a
web server which should be reachable from your browser. The authorization flow
will first direct you to the Fitbit web page which will redirect you back to
this application's webserver.

**Note**: this application's webserver is _NOT_ the same as the webserver of
your Workout Tracker!

This linking step is only needed the first time, when granting access to your
account. After the link is done, the application will store the refresh token
and use this to continue syncing activities.

### Prepare a redirect URI

The redirect URI is a full URL (including protocol) which your browser will be
redirected to to finalize the OAuth flow. The URI should therefore be reachable
from your browser, not from the Internet.

This application will, by default bind to `http://localhost:8080` and respond to
`/link`. If you run the application on the same machine as your browser, the URI
might be <http://localhost:8080/link>.

### Detailed steps

- Log in on the developer portal: <https://dev.fitbit.com/apps>
- Register a new app: <https://dev.fitbit.com/apps/new>
- Fill in the form:
  - **Application Name**: choose for yourself, eg. "Workout Tracker"
  - **Description**: choose for yourself
  - **Application Website URL**: I used the root url of my Workout Tracker
    installation; this is only for documentation purposes, since the cli will
    run separately from the Workout Tracker webserver.
  - **Organization**: choose for yourserlf, eg. "Home"
  - **Organization Website URL**: I used the same as the root url
  - **Terms of Service URL**: I used the same as the root url
  - **Privacy Policy URL**: I used the same as the root url
  - **OAuth 2.0 Application Type**: Server
  - **Redirect URL**: the redirect URI as defined above
  - **Default Access Type**: "Read Only"

After filling this form, you should receive the following information:

- Client ID
- Client Secret

In order to complete the link, you also need your own Fitbit user ID. You can
find this when logging in to the regular
[Fitbit dashboard](https://www.fitbit.com/dashboard) and clicking on your avatar
in the top right corner. The new URL should be something like
<https://www.fitbit.com/user/XYZ123>, where `XYZ123` is your user ID.

Now you can use this information to link the application:

```bash
fitbit-sync link --client-id CLIENT_ID --client-secret CLIENT_SECRET --user-id USER_ID
```

The application will create a configuration file with the above information and
store it in the default user configuration directory (eg. on Linux:
`$HOME/.config/workout-tracker/fitbit.json`).
