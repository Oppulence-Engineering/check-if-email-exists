# Changelog

## [1.0.0](https://github.com/Oppulence-Engineering/check-if-email-exists/compare/v0.11.7...v1.0.0) (2025-11-28)


### ⚠ BREAKING CHANGES

* 
* **core:** 
* 
* **core:** 
* 
* **backend:** 
* Rename all VerifyMethod to VerifMethod
* For Hotmail, Gmail and Yahoo addresses, the `*_use_api` and `*_use_headless` parameters have been removed and replaced with a `*VerifyMethod`, an enum which can take value Api, Headless or Smtp. If using headless, pass a webdriver address to env variable RCH_WEBDRIVER_ADDR.
* `input.hotmail_use_headless` is now a bool instead of a string. Pass the webdriver address as an environment variable `RCH_WEBDRIVER_ADDR` now.
* 
* 

### Features

* **#289:** add haveibeenpwned check ([#1253](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1253)) ([166dbd2](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/166dbd2cc878e30c51538b919abc1aaea4465c45))
* Add `/v1/{check_email,bulk}` endpoints with throttle&concurrency ([#1537](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1537)) ([08522e4](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/08522e4326bbcbc980cf501d5d994d0c17222561))
* Add `misc.is_b2c` field ([#1553](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1553)) ([14a6759](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/14a6759d805d2051a4a1e1d81588279cb9c85336))
* Add AWS SQS support ([#1554](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1554)) ([92be54e](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/92be54ebfe4a2d19101141f55e94fc8e9588ff95))
* Add back RabbitMQ-based worker ([#1513](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1513)) ([de75ece](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/de75eceef32c6ea512e0a301ec62d393bb59ff0f))
* Add debug information about each email verification ([#1391](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1391)) ([3ea6e66](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/3ea6e6607735682dfca6ecfa27460650ac6e42d3))
* add email address normalisation ([#1206](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1206)) ([f8ec348](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/f8ec348883cd4f4a20a8acbb38d54b69e798222b))
* Add optional timeout on proxy (env var: `RCH__PROXY__TIMEOUT_MS`) ([#1595](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1595)) ([0e51eb6](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/0e51eb686dad6bd2ec827e785bf9c30ccc88cde1))
* Add RabbitMQ worker ([#1395](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1395)) ([ecef8c9](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/ecef8c98deb744390c7017a4e98d4f3c7e737fcb))
* Add suggestions for syntax errors ([#1192](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1192)) ([2d385f3](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/2d385f30f7a62ab2706599fbb89fb50275cffb5f))
* Allow /v1/check_email without worker mode ([9ca9f39](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/9ca9f39ee487dc1b7d9b4cdc9a0b2c0669b10bc0))
* Allow multiple proxies ([#1562](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1562)) ([eed5a15](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/eed5a1536af37877f12eebab6481acaa6efa55c5))
* **backend:** Add header secret to protect against public requests ([#1158](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1158)) ([fa6a56b](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/fa6a56b62f4b3aeeec704cfe4882755998d40833))
* **backend:** Add one simple retry on Unknown ([fcffc1a](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/fcffc1a28bab990b0596ad8b66163e47a494191b))
* **backend:** Add POST /v1/bulk ([#1413](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1413)) ([d9302d4](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/d9302d4c1cec6a5a1788afe2a3718df8986f118f))
* **backend:** Add reply-to queue ([aaea59f](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/aaea59f251634db7c35f029b09ef6e5f8c77cfbc))
* **backend:** Add worker webhook ([db90cfa](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/db90cfa27b85916685268a3599bdfdb2c46de07a))
* **backend:** Customize SMTP defaults ([8f152b8](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/8f152b83c70b94618b71308552a6999f4b27aa2f))
* **backend:** Prune bulk email verification database ([#1377](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1377)) ([f905735](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/f90573566abf40133ebfb28ebc8f18ad8278a9b3))
* **backend:** Reject a request with to_email field empty or missing ([#1353](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1353)) ([1d9c29f](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/1d9c29f5a48655a11f985b7df91c8bcbdf102487))
* **backend:** Remove /v0/bulk endpoints ([#1421](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1421)) ([522f324](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/522f32448416cd75a70ddb51038e50d06c3130b4))
* **backend:** Support RCH_SMTP_TIMEOUT ([#1407](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1407)) ([b9bda40](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/b9bda4049540372811a86d8dd7ba873c9875e54d))
* **core:** Add check gravatar image ([#1188](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1188)) ([6a26035](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/6a26035327ab681a65a4f4ba284e155f00680e89))
* **core:** Add Hotmail checks via headless password recovery ([#1165](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1165)) ([7517ed9](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/7517ed98ba966158deebba6a1a4745c931bfed18))
* **core:** Update async-smtp to 0.9 ([#1520](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1520)) ([297ce4f](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/297ce4f11994b483faa015bebe4abf550eb77e11))
* Increase content length limit for bulk validation endpoint ([#1525](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1525)) ([bbdab31](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/bbdab31e0dde54d21f4eeb5880ae28e60de7dced))
* Move `backend` code to this repo ([#1138](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1138)) ([0dc6053](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/0dc60531d26efb217137347ef2b6aaf678d94238))
* Revert back to `check_email` input with single email ([#1150](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1150)) ([ce1ba53](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/ce1ba5346849b578a0ed30b1d72096f15cfbc09d))
* Set default timeout to 10s ([#1251](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1251)) ([d04f84c](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/d04f84cc1e7b30e02d3717ab1af9f680cdb2c27f))
* Yahoo account recovery via headless ([#1364](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1364)) ([6f0f12b](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/6f0f12b8cf528e819f8743f7e3c5f5e141c51559))


### Bug Fixes

* Add "utilisateur inconnu" in invalid parser ([#1594](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1594)) ([fb91653](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/fb9165303e2d7be59ed2fa4f0682e8592bc0c5e7))
* Add backend_name in /v0/check_email ([a738fae](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/a738faec99942d20b817298f7850e84ab3e74835))
* **backend:** CSV download retrieves all results ([#1362](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1362)) ([b3670fc](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/b3670fcaebce05a0aab09bcc3253134cb3c643c1))
* **backend:** Fix docker CTRL+C ([3a7245f](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/3a7245f9a47e8332d682d437d9492559e5adf66f))
* **backend:** Fix dockerfile ([f0ed49f](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/f0ed49f50238c1c71a130f3db19ec047af00b8df))
* **backend:** Fix env var for multiple queues ([ed19166](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/ed191662b18c62f397b4fed6b95249b5aa76c423))
* **backend:** Improve sentry error messages ([#1155](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1155)) ([d90d998](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/d90d998d1cb189fed3f888659aa08fd4fabf6e93))
* **backend:** Redact email in sentry bug tracking ([2c2d1d8](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/2c2d1d88c0086196bc09359e32c96638124d9539))
* **backend:** Update sqlx to 0.7 ([#1390](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1390)) ([7198f87](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/7198f87de92ab403cdc1e7c68667cdef9db96085))
* Bring back `{yahoo,hotmailb2c}_verif_method` ([#1606](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1606)) ([3fbe520](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/3fbe5200a3d8608fbd72c0f2a5917326c1f8ec91))
* **core:** Clean up CheckEmailInput ([#1531](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1531)) ([b97b9ff](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/b97b9ff9b91bdfbf18e5c0892559e87e7cd5e16c))
* **core:** Fix default CheckEmailInput ([09215a1](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/09215a13ac3525861e6cd1dea3fc71c13dfffe52))
* **core:** Fix MX random record selection ([#1263](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1263)) ([9fae593](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/9fae593b8590ad5efb3e7d16bbd25cc05c228cb9))
* **core:** Headless check for Microsoft365 too ([#1346](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1346)) ([682cc2d](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/682cc2d96b93d73f3fca3ba11f03800477c8fb9e))
* **core:** Improve invalid parser ([#1166](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1166)) ([bb46004](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/bb460046bf1cb031fee706d836c8a737157f803c))
* **core:** Use Smtp for Gmail by default ([8e79884](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/8e79884314f0c1eec5a7964fa686e2c60e7d2209))
* **core:** Use tagged enum representation ([ffde851](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/ffde851068798adc3372d843a916a121b5caeccb))
* **docker:** Fix dockerfile entrypoint ([d1d3326](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/d1d3326af88a85b2192796d8d2c92ff854b5644d))
* Fix dockerfile ([ce5067e](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/ce5067e4050e0cf3fa6c022bc7e25e5f15261c2a))
* Fix rabbitmq docker compose ([7c3856e](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/7c3856ebec6089b37b3dd30e3c4f13df9fb4e73a))
* Fix version in logs ([fa6be78](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/fa6be7867abae981b0d82fde24e0310b9759ab1f))
* Improve logging, add retries for Yahoo headless, switch to rustls ([#1549](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1549)) ([b1377db](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/b1377db2b32155d766a09a76864fc9b0990833e6))
* Make new config backwards-compatible ([#1567](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1567)) ([b824e2c](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/b824e2c988ee4eef021b97fc65ebcfa36a166d7f))
* Put Smtp debug details in Debug struct ([5b71ca5](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/5b71ca59b6fab18263348aeafc7a895b7f4b8076))
* Reinstate proxy in JSON request ([#1569](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1569)) ([c36e6e0](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/c36e6e09c9079de210d288b84d79b984e2ea77f0))
* Remove local_ip retrieval ([ff8e599](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/ff8e5998f8b88954b4104f9251d1331542dbb182))
* Remove max requests per minute/day ([07a6d96](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/07a6d96416f52ac0824e7e7ac665fd2169ddc7ec))
* Revert back to using lowest-priority MX record ([#1578](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1578)) ([60468b3](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/60468b3f533491a0dff6a42e7096f34ece19896c))
* Show thread ID in logs ([#1579](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1579)) ([3388163](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/3388163d03b66ba92455be8404441e8555a9d53c))
* Support queues in env var ([39655d5](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/39655d51afe5f65d62cd5dc3485586e16bcdec31))
* Typo in expect of RCH_VERIF_METHOD ([#1405](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1405)) ([c50d8eb](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/c50d8ebdfc470fe1ec6290e07668c70095298799))
* Use chromedriver instead of gecko for parallel requests ([e282e28](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/e282e28aeb7259d800f7faad97173c3a216095a4))


### Reverts

* "Show thread ID in logs ([#1579](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1579))" ([56e7838](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/56e7838f28067b05b58f1fcd166368a915aafbbc))
* **backend:** Bring back the sqlxmq-based bulk verification ([#1477](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1477)) ([322ad4e](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/322ad4e4b53d534a8ae6461f3d3383d67b219b5d))


### Miscellaneous Chores

* Rename all VerifyMethod to VerifMethod ([9f9607d](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/9f9607d35478a1051dde56812f8914ff75d4c5ac))


### Code Refactoring

* Change RUST_LOG target to `reacher` ([#1152](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1152)) ([7e87be2](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/7e87be26f1e35a6936bfc967c872cd42b93fd256))
* Use config-rs instead of env vars ([#1530](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1530)) ([bcd2dc8](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/bcd2dc867b7dc2bdaeb70097fd14109c2a40da17))
* Use verify method for known providers ([#1366](https://github.com/Oppulence-Engineering/check-if-email-exists/issues/1366)) ([5ca4dfa](https://github.com/Oppulence-Engineering/check-if-email-exists/commit/5ca4dfa5ec38fba0ec7cfb052106da8d6af4df44))
