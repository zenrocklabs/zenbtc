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
