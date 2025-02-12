## [1.12.6](https://github.com/zenrocklabs/zenbtc/compare/v1.12.5...v1.12.6) (2025-02-12)


### Bug Fixes

* add store getters ([85b65fc](https://github.com/zenrocklabs/zenbtc/commit/85b65fc47d7ffc1ce3991b15e0c9bb46c5fa24e8))

## [1.12.5](https://github.com/zenrocklabs/zenbtc/compare/v1.12.4...v1.12.5) (2025-02-12)


### Bug Fixes

* add WalkBurnEvents method ([f70ffc1](https://github.com/zenrocklabs/zenbtc/commit/f70ffc1f14f88266e4c0b8234ddbfceee47f017d))

## [1.12.4](https://github.com/zenrocklabs/zenbtc/compare/v1.12.3...v1.12.4) (2025-02-12)


### Bug Fixes

* minor polish ([88a8938](https://github.com/zenrocklabs/zenbtc/commit/88a89387550d23a506ec983faf9dea7d2d87cc5e))

## [1.12.3](https://github.com/zenrocklabs/zenbtc/compare/v1.12.2...v1.12.3) (2025-02-12)


### Bug Fixes

* add CreateBurnEvent method ([33d7535](https://github.com/zenrocklabs/zenbtc/commit/33d7535161d332385551b64bb3af752856687509))

## [1.12.2](https://github.com/zenrocklabs/zenbtc/compare/v1.12.1...v1.12.2) (2025-02-12)


### Bug Fixes

* update BurnEvents to use map + improve queries ([#30](https://github.com/zenrocklabs/zenbtc/issues/30)) ([a159929](https://github.com/zenrocklabs/zenbtc/commit/a1599291ff8f087d447c390165969fa5e3e96239))

## [1.12.1](https://github.com/zenrocklabs/zenbtc/compare/v1.12.0...v1.12.1) (2025-02-12)


### Bug Fixes

* add check to ensure key has metadata and is therefore a ZenBTC key ([674d92d](https://github.com/zenrocklabs/zenbtc/commit/674d92dcc55f238c6fe89734a93c539edfe4775a))
* re-order metadata test before supply change, and check for nil metadata ([4cb8ff4](https://github.com/zenrocklabs/zenbtc/commit/4cb8ff403caf7c3cffec95ffa51a971d976a07fb))

# [1.12.0](https://github.com/zenrocklabs/zenbtc/compare/v1.11.2...v1.12.0) (2025-02-10)


### Features

* store redemption txbytes in sign request ([a42e275](https://github.com/zenrocklabs/zenbtc/commit/a42e275633a3c117561ca35230bc69ba5c9bc43a))

## [1.11.2](https://github.com/zenrocklabs/zenbtc/compare/v1.11.1...v1.11.2) (2025-02-07)


### Bug Fixes

* add `EthTokenAddr` param field ([c363d76](https://github.com/zenrocklabs/zenbtc/commit/c363d76f629be1264f8131bda95cbe4745a2fa9d))

## [1.11.1](https://github.com/zenrocklabs/zenbtc/compare/v1.11.0...v1.11.1) (2025-02-07)


### Bug Fixes

* `VerifyDepositBlockInclusion` handler ([3be9dbf](https://github.com/zenrocklabs/zenbtc/commit/3be9dbf5b15fe2f5b85692d25a2dfa0a455be9a1))

# [1.11.0](https://github.com/zenrocklabs/zenbtc/compare/v1.10.0...v1.11.0) (2025-02-06)


### Features

* multichain impl ([#27](https://github.com/zenrocklabs/zenbtc/issues/27)) ([e7de42e](https://github.com/zenrocklabs/zenbtc/commit/e7de42e058562aa0e39aa0958a37caefb199a7df))

# [1.10.0](https://github.com/zenrocklabs/zenbtc/compare/v1.9.6...v1.10.0) (2025-02-06)


### Features

* Add CAIP2 functionality ([#28](https://github.com/zenrocklabs/zenbtc/issues/28)) ([bb725e7](https://github.com/zenrocklabs/zenbtc/commit/bb725e78035e561a0f2d737c2f42614af8e4fe84))

## [1.9.6](https://github.com/zenrocklabs/zenbtc/compare/v1.9.5...v1.9.6) (2025-02-03)


### Bug Fixes

* add temp debug log ([3df533d](https://github.com/zenrocklabs/zenbtc/commit/3df533d0c166dc3313b8021dde764849c5e06fd3))

## [1.9.5](https://github.com/zenrocklabs/zenbtc/compare/v1.9.4...v1.9.5) (2025-02-03)


### Bug Fixes

* improve supply query ([de371dc](https://github.com/zenrocklabs/zenbtc/commit/de371dc19e81b972524e85d16d77624ff347cd84))

## [1.9.4](https://github.com/zenrocklabs/zenbtc/compare/v1.9.3...v1.9.4) (2025-02-03)


### Bug Fixes

* supply query ([2762766](https://github.com/zenrocklabs/zenbtc/commit/27627665cc801171496c8cf3f29bfd47ede2cb4b))

## [1.9.3](https://github.com/zenrocklabs/zenbtc/compare/v1.9.2...v1.9.3) (2025-01-30)


### Bug Fixes

* migration test import ([9b0f326](https://github.com/zenrocklabs/zenbtc/commit/9b0f3266b94b7ec4b70ee1f13237382ef3f43489))

## [1.9.2](https://github.com/zenrocklabs/zenbtc/compare/v1.9.1...v1.9.2) (2025-01-30)


### Bug Fixes

* revert caip2 changes + add migration to fix state ([70024c3](https://github.com/zenrocklabs/zenbtc/commit/70024c34c58a4fad353d4afb1fe60b6556e84b14))

## [1.9.1](https://github.com/zenrocklabs/zenbtc/compare/v1.9.0...v1.9.1) (2025-01-28)


### Bug Fixes

* permit redemptions with no outputs to enable consolidation reansactions ([6a52f3b](https://github.com/zenrocklabs/zenbtc/commit/6a52f3b649ac36f77a7db81f24b3fec445fee8ff))

# [1.9.0](https://github.com/zenrocklabs/zenbtc/compare/v1.8.0...v1.9.0) (2025-01-28)


### Features

* Add caip2 id field ([#22](https://github.com/zenrocklabs/zenbtc/issues/22)) ([319c7dc](https://github.com/zenrocklabs/zenbtc/commit/319c7dc07e89629da339cb1cbc6458a8cfd88fad))

# [1.8.0](https://github.com/zenrocklabs/zenbtc/compare/v1.7.17...v1.8.0) (2025-01-27)


### Features

* create README.md ([#24](https://github.com/zenrocklabs/zenbtc/issues/24)) ([a56c783](https://github.com/zenrocklabs/zenbtc/commit/a56c783c16c37dbeaeba13dc8046e07eb5ffe3cd))

## [1.7.17](https://github.com/zenrocklabs/zenbtc/compare/v1.7.16...v1.7.17) (2025-01-20)


### Bug Fixes

* store migration keyIDs ([f011e3b](https://github.com/zenrocklabs/zenbtc/commit/f011e3b4bb80a85d9e566ff8230c45818195df18))

## [1.7.16](https://github.com/zenrocklabs/zenbtc/compare/v1.7.15...v1.7.16) (2025-01-17)


### Bug Fixes

* bump zrchain ver, revert cosmos sdk fork ver, remove extraneous migration files ([42a3576](https://github.com/zenrocklabs/zenbtc/commit/42a35768861bb2978d864020f36eca968ccff525))

## [1.7.15](https://github.com/zenrocklabs/zenbtc/compare/v1.7.14...v1.7.15) (2025-01-16)


### Bug Fixes

* track `PendingZenBTC` supply to use in exchange rate calc ([c66d066](https://github.com/zenrocklabs/zenbtc/commit/c66d066f53aa25049e325ce6c48e242d9092f275))

## [1.7.14](https://github.com/zenrocklabs/zenbtc/compare/v1.7.13...v1.7.14) (2025-01-15)


### Bug Fixes

* add `PrevNonce` field ([761a396](https://github.com/zenrocklabs/zenbtc/commit/761a396237acb17608fcd92830add309600bf52b))

## [1.7.13](https://github.com/zenrocklabs/zenbtc/compare/v1.7.12...v1.7.13) (2025-01-14)


### Bug Fixes

* temporarily hardcode bitcoin proxy address ([66c681d](https://github.com/zenrocklabs/zenbtc/commit/66c681d549de1685384e0b37d9ff818839db67ec))

## [1.7.12](https://github.com/zenrocklabs/zenbtc/compare/v1.7.11...v1.7.12) (2025-01-14)


### Bug Fixes

* add logging (params issue) ([d9f8dc3](https://github.com/zenrocklabs/zenbtc/commit/d9f8dc3be4117ec5c228128b7ea4a61fdc5fa512))

## [1.7.11](https://github.com/zenrocklabs/zenbtc/compare/v1.7.10...v1.7.11) (2025-01-14)


### Bug Fixes

* exchange rate calculation ([8a78b60](https://github.com/zenrocklabs/zenbtc/commit/8a78b602cd3ad755d7f17b91de04e32e2e1e5381))

## [1.7.10](https://github.com/zenrocklabs/zenbtc/compare/v1.7.9...v1.7.10) (2025-01-14)


### Bug Fixes

* exchange rate calculation ([d701de4](https://github.com/zenrocklabs/zenbtc/commit/d701de46a90f0345b100e28c6f990b3fcb22e840))

## [1.7.9](https://github.com/zenrocklabs/zenbtc/compare/v1.7.8...v1.7.9) (2025-01-13)


### Bug Fixes

* enable `UpdateParams` CLI cmd + bump cosmos fork ver ([#21](https://github.com/zenrocklabs/zenbtc/issues/21)) ([daa7ac6](https://github.com/zenrocklabs/zenbtc/commit/daa7ac6353ed3e7e8ef8afe8644402d17e9e5384))

## [1.7.8](https://github.com/zenrocklabs/zenbtc/compare/v1.7.7...v1.7.8) (2025-01-13)


### Bug Fixes

* params migration ([#20](https://github.com/zenrocklabs/zenbtc/issues/20)) ([ac4bf49](https://github.com/zenrocklabs/zenbtc/commit/ac4bf49ef2602b3126c005d9453e331889f9fe14))

## [1.7.7](https://github.com/zenrocklabs/zenbtc/compare/v1.7.6...v1.7.7) (2025-01-12)


### Bug Fixes

* migration version ([#19](https://github.com/zenrocklabs/zenbtc/issues/19)) ([d668796](https://github.com/zenrocklabs/zenbtc/commit/d668796df7ecd7e87e594d5b31c318d31ea8fb74))

## [1.7.6](https://github.com/zenrocklabs/zenbtc/compare/v1.7.5...v1.7.6) (2025-01-10)


### Bug Fixes

* migration files ([#18](https://github.com/zenrocklabs/zenbtc/issues/18)) ([07ff35f](https://github.com/zenrocklabs/zenbtc/commit/07ff35f09f2c0c7270d17b66230ba331a799a6a4))

## [1.7.5](https://github.com/zenrocklabs/zenbtc/compare/v1.7.4...v1.7.5) (2025-01-10)


### Bug Fixes

* move most zenBTC logic to `x/zenbtc` module + add vout key field `LockTransactions` ([#17](https://github.com/zenrocklabs/zenbtc/issues/17)) ([bb9b8e8](https://github.com/zenrocklabs/zenbtc/commit/bb9b8e82ad8b476dae094f6356ceb9084287cf5b))

## [1.7.4](https://github.com/zenrocklabs/zenbtc/compare/v1.7.3...v1.7.4) (2025-01-09)


### Bug Fixes

* add logging to debug mint issue ([#16](https://github.com/zenrocklabs/zenbtc/issues/16)) ([68fbb2a](https://github.com/zenrocklabs/zenbtc/commit/68fbb2aa5996d9729c6e504f6c985007ca3639dd))

## [1.7.3](https://github.com/zenrocklabs/zenbtc/compare/v1.7.2...v1.7.3) (2025-01-07)


### Bug Fixes

* query unstaked redemptions ([#15](https://github.com/zenrocklabs/zenbtc/issues/15)) ([a2ce053](https://github.com/zenrocklabs/zenbtc/commit/a2ce053798e6e68ac3b1854802c6e032e2db53e4))

## [1.7.2](https://github.com/zenrocklabs/zenbtc/compare/v1.7.1...v1.7.2) (2024-12-31)


### Bug Fixes

* KeyByAddress, wallet type to specify the correct type based on the BTC Chain in use. ([#14](https://github.com/zenrocklabs/zenbtc/issues/14)) ([9d829a1](https://github.com/zenrocklabs/zenbtc/commit/9d829a1d8b0fc1487984c0120ec06592893ae614))

## [1.7.1](https://github.com/zenrocklabs/zenbtc/compare/v1.7.0...v1.7.1) (2024-12-30)


### Bug Fixes

* remove BTC Proxy debugging code for block generated before zrchain start ([#13](https://github.com/zenrocklabs/zenbtc/issues/13)) ([53ad7f0](https://github.com/zenrocklabs/zenbtc/commit/53ad7f02c81ba910b7eeb6bdff7d3a24af66621e))

# [1.7.0](https://github.com/zenrocklabs/zenbtc/compare/v1.6.2...v1.7.0) (2024-12-23)


### Features

* unwrap flow + fixes ([#12](https://github.com/zenrocklabs/zenbtc/issues/12)) ([d0d44b7](https://github.com/zenrocklabs/zenbtc/commit/d0d44b731ca9f55fc6e41651109f6ae5356d5179))

## [1.6.2](https://github.com/zenrocklabs/zenbtc/compare/v1.6.1...v1.6.2) (2024-12-18)


### Bug Fixes

* bump zrchain ver ([46d47e3](https://github.com/zenrocklabs/zenbtc/commit/46d47e37f1a75492d16420997ac28bcfae9d8300))

## [1.6.1](https://github.com/zenrocklabs/zenbtc/compare/v1.6.0...v1.6.1) (2024-12-18)


### Bug Fixes

* revert temp commit ([#11](https://github.com/zenrocklabs/zenbtc/issues/11)) ([646220d](https://github.com/zenrocklabs/zenbtc/commit/646220d112cfbdc0cde6217aa0a58515f5d1e4c6))

# [1.6.0](https://github.com/zenrocklabs/zenbtc/compare/v1.5.1...v1.6.0) (2024-12-18)


### Bug Fixes

* temporarily comment new zrchain code ([6aec17b](https://github.com/zenrocklabs/zenbtc/commit/6aec17bd108f295181715100921c16031f4ea564))
* unhappy path `VerifyDepositBlockInclusion` ([d4cdcd9](https://github.com/zenrocklabs/zenbtc/commit/d4cdcd90d510737d83c558ef72c11682184074b9))


### Features

* add redemptions query ([af30094](https://github.com/zenrocklabs/zenbtc/commit/af300946b6484884728b2539348442b2650f023f))
* add rewards deposit key, bypass minting on rewards deposit ([587bff1](https://github.com/zenrocklabs/zenbtc/commit/587bff144ef64fb7db733c05f000d16efcc643e5))
* add tx handler for `SubmitUnsignedRedemptionTx` ([1190f7c](https://github.com/zenrocklabs/zenbtc/commit/1190f7ce38efdbdb58c75b61dc2adda4acd63e47))
* unwrap + yield flow wip (w/ temp commented code) ([6da39df](https://github.com/zenrocklabs/zenbtc/commit/6da39df0262f90335c4ab2cf0adcb4be3a3e9743))
* zenBTC exchange rate ([45b1416](https://github.com/zenrocklabs/zenbtc/commit/45b14165f80f98ed9634077cd8c08527ca6036e8))

## [1.5.1](https://github.com/zenrocklabs/zenbtc/compare/v1.5.0...v1.5.1) (2024-12-02)


### Bug Fixes

* remove obsolete proto ([8e86adb](https://github.com/zenrocklabs/zenbtc/commit/8e86adb40d73d143de41853d1895ca467aaac1c0))
* remove obsolete proto ([8bf062f](https://github.com/zenrocklabs/zenbtc/commit/8bf062f5be13add3fd04f96000f4bdb2d8559726))

# [1.5.0](https://github.com/zenrocklabs/zenbtc/compare/v1.4.2...v1.5.0) (2024-11-29)


### Bug Fixes

* get converter running + identify failure source ([e948360](https://github.com/zenrocklabs/zenbtc/commit/e948360ab382998b05f0ff334a398dea5cc3734d))


### Features

* implement working thorchain swaps ([a621592](https://github.com/zenrocklabs/zenbtc/commit/a6215923bf059117442c17ccdbcda7967c393087))
* THORChain Rewards Conversion ([240294e](https://github.com/zenrocklabs/zenbtc/commit/240294e8144d31f669d514d4c8a4acafdd75ee42))

## [1.4.2](https://github.com/zenrocklabs/zenbtc/compare/v1.4.1...v1.4.2) (2024-11-13)


### Bug Fixes

* update proto for redemptions ([0f4b8fa](https://github.com/zenrocklabs/zenbtc/commit/0f4b8fa73203459f01e8675c1b1f239ce6cf1249))

## [1.4.1](https://github.com/zenrocklabs/zenbtc/compare/v1.4.0...v1.4.1) (2024-11-13)


### Bug Fixes

* add ids to redemptions ([f4c62cf](https://github.com/zenrocklabs/zenbtc/commit/f4c62cf89e4e95df5f116fa7ab10cc1342895c39))

# [1.4.0](https://github.com/zenrocklabs/zenbtc/compare/v1.3.1...v1.4.0) (2024-11-13)


### Features

* redemption types ([e81d8e3](https://github.com/zenrocklabs/zenbtc/commit/e81d8e3c575209c462e7782922b4c34f79264554))

## [1.3.1](https://github.com/zenrocklabs/zenbtc/compare/v1.3.0...v1.3.1) (2024-11-12)


### Bug Fixes

* update `RedemptionTracker` binding ([b0ab8b9](https://github.com/zenrocklabs/zenbtc/commit/b0ab8b94278e74addd1161054a2a41adfec443ab))

# [1.3.0](https://github.com/zenrocklabs/zenbtc/compare/v1.2.1...v1.3.0) (2024-11-12)


### Features

* add `RedemptionTracker` binding ([74a1adb](https://github.com/zenrocklabs/zenbtc/commit/74a1adbe53c8a3f804b242d314815f77658b3de4))

## [1.2.1](https://github.com/zenrocklabs/zenbtc/compare/v1.2.0...v1.2.1) (2024-11-07)


### Bug Fixes

* override ([88e80cf](https://github.com/zenrocklabs/zenbtc/commit/88e80cff8e3656a449e76fc4a99bb11182ce1049))
* pulsar generation ([755b6c6](https://github.com/zenrocklabs/zenbtc/commit/755b6c6dd9dc04365f060cb42f40b2db983852c7))

# [1.2.0](https://github.com/zenrocklabs/zenbtc/compare/v1.1.0...v1.2.0) (2024-11-06)


### Bug Fixes

* generate pulsar files ([5a3379c](https://github.com/zenrocklabs/zenbtc/commit/5a3379cb7e5e5a6555380df0804f49e5c3f1a360))


### Features

* better error handling and performance ([67383da](https://github.com/zenrocklabs/zenbtc/commit/67383dad7d6bb85a7ad2dd82b20142b649d7042b))

# [1.1.0](https://github.com/zenrocklabs/zenbtc/compare/v1.0.3...v1.1.0) (2024-11-06)


### Bug Fixes

* protogen to generate pb and python files ([181eded](https://github.com/zenrocklabs/zenbtc/commit/181eded2bf658838758bbcd82e7ff9b3f501b1a6))


### Features

* semantic release setup ([55db4d7](https://github.com/zenrocklabs/zenbtc/commit/55db4d7d64899fc7e3312927e4b209135d3ce9dc))
* semantic release setup ([a5b4c45](https://github.com/zenrocklabs/zenbtc/commit/a5b4c456f9a00f100201f92ad6706b112ab9da2a))
