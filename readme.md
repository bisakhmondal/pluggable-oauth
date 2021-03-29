## The project is created to demonstrate how the OAuth2.0 authentication flow is going to be implemented (as a proposal).

- Though there are several identity providers to authenticate users using OAuth along with OpenID Connect, the project fully respects the admin's preferences of choosing one provider over the another.
- Currently, this project supports authentication using `Google`, `Facebook`, `Github`, but the way the project has been developed, it provides a fully flexible way to incorporate multiple identity providers in future (if necessary).
- As the preference of the admin is the top priority, there may be cases when one or two or no oauth identity provider has been chosen by the admin. `So instead of polluting the UI with the icons of all supported identity providers, it provides a kind of pluggable feature to add or remove the icons from the login page by simple editing of JSON for the frontend.`

The UI / Login screen only shows the icons for OAuth for which an entry has been made in the [oauth-config.json](./go-auth-frontend/src/oauth-config.json)
for eg.
```json
{
  "provider":"Google",
  "backend_url": "google/login",
  "icon": "https://cdn.pixabay.com/photo/2015/12/11/11/43/google-1088004_1280.png"
}
```
If the entry has been removed from the json, there will be no option in the dashboard to `login using google`.


- In the backend, separates routes has been created for authentication using different providers to fetch the `email` scope that we are interested in for authentication.
  (It is required because different OAuth provider authorization server & resource server grants userinfo from accessToken is different).
  
- Now to enable the authentication using a identity provider (for say Google) the admin has to generate a `client_id` & `client_secret` from the designated source(for say GCP). And just have to put that into `yaml` file.

- A yaml file [conf.yaml](./go-oauth-backend/conf.yml) has been created where the admin will just provide the `client_id` & `client_secret` reveived from the provider. The server will take care rest of it.
an example
```yaml
credentials:
  google:
    id: "569278128850-e120l986lf4ck9i2ne4mv4s7edop8l53.apps.googleusercontent.com"
    secret: "DNzgA9FcTGQ61FGsAes5D6Kv"

  facebook:
    id: "781905489374910"
    secret: "f0179945d2d929e7c93b0c68bb4039d7"


```

- **So in brief, suppose I am an admin and I want to incorporate `login with google` in the dashboard i have to do two things.**
- **For frontend rendering of google logo, I will create an entry in mentioned JSON file.**
- **For backend authentication process, I will put the `id` & `secret` in the yaml file.**

And, that's it.

For demonstration purpose, on a good faith I have provided with the id & secrets for google, facebook & github.

Just do the following
```shell
# run the react UI
cd go-auth-frontend 
npm i
npm start # make sure it starts the UI at port 3000

# run the backend
cd  go-oauth-backend
go get -u 
go run main.go parse.go # make sure it starts the backend at port 4000
```

Thank you.
