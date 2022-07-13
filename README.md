# TODO
- Parse Args
  - [x] Run all pre add arg hooks
  - [x] Run all add arg hooks
  - [ ] Handle anything that is needed with the args
  - [x] Run all post add arg hooks
- Get Profiles(args)
  - [x] Get AWS files (config and credentials)
  - [x] Run all pre collect profiles hooks
  - [x] Run all collect profiles hooks
  - [x] Aggregate profiles
  - [x] Run all post collect profiles hooks
- Get Credentials(args, profiles)
  - [ ] Run all pre get creds
  - [ ] Check args for auto_refresh and arg.json, with_saml, with_web_identity, or else
    - [ ] Pull credentials from 'Credentials'
    - [ ] Run get saml credentials
    - [ ] Run all get credentials with web identity hooks
    - [ ] Run all get credentials hooks
    - [ ] Figure out stuff for autoawsume
    - [ ] Run all post get credentials hooks
    - [ ] Handle if we don't have credentials yet at this point
- [ ] Export credentials data


- Look at default plugins and create a plugin for it
