# Verthash One-Click Miner

The Verthash One Click Miner is functionally the same as [upstream](https://github.com/vertcoin-project/one-click-miner-vnext) utilizing a new data directory, `verthash-ocm`.

This allows the default [Zergpool](https://zergpool.com/site/faq) payout in VTC to be changed to a user-inputted Dogecoin, Bitcoin, Litecoin, Bitcoin Cash or Dash address. Pending payout information will accurately track the payout of choice although Expected Earnings (24h) still reflects an amount in VTC.


**Current Zergpool payout thresholds as of February 18, 2021:**

`DOGE payment threshold is 400 DOGE and 80 DOGE on Sunday late evening(CET)`

`BTC payment threshold is 0.0025 BTC`

`LTC payment threshold is 0.25 LTC and 0.05 LTC on Sunday late evening(CET)`

`BCH payment threshold is ~0.00712 and ~0.00142 on Sunday late evening(CET)`

`DASH payment threshold is ~0.0169 and ~0.00338 on Sunday late evening(CET)`

*Network and exchange fees may apply

The One-Click Miner allows cryptocurrency enthusiasts to get into mining with minimal effort. When you download the One-Click Miner, you will be asked to provide a password for your (built in) wallet. It will then immediately commence mining.

This software is available for Windows and Linux.

## FAQ

### Which GPUs are supported?

Please refer to this list of [supported hardware.](https://github.com/CryptoGraphics/VerthashMiner#supported-hardware)

### I have an error message that reads 'Failure to configure'

You may need to add an exclusion to your antivirus / Windows Defender.

### My GPU is supported but an error messages reads 'no compatible GPUs'

Update your GPU drivers to the latest version.


## Building

The GUI of this MVP is based on [Wails](https://wails.app) and [Go](https://golang.org/).

Install the Wails [prerequisites](https://wails.app/home.html#prerequisites) for your platform, and then run:

```bash
go get github.com/wailsapp/wails/cmd/wails
```

Then clone this repository, and inside its main folder, execute:

```bash
wails build
```

## Donations

If you want to support the further development of the upstream One Click Miner, feel free to donate Vertcoin to [Vmnbtn5nnNbs1otuYa2LGBtEyFuarFY1f8](https://insight.vertcoin.org/address/Vmnbtn5nnNbs1otuYa2LGBtEyFuarFY1f8).
