import { GOOGLE_OAUTH_CLIENT_ID, REDIRECT_URL } from "../constants";

export const getGoogleUrl = (state: string) => {
  const rootUrl = `https://accounts.google.com/o/oauth2/v2/auth`;

  const options = {
    redirect_uri: REDIRECT_URL,
    client_id: GOOGLE_OAUTH_CLIENT_ID,
    access_type: "offline",
    response_type: "code",
    prompt: "consent",
    scope: [
      "https://www.googleapis.com/auth/userinfo.email",
    ].join(" "),
    state: state,
  };

  const qs = new URLSearchParams(options);

  return `${rootUrl}?${qs.toString()}`;
};
