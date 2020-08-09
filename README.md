# Forex Exporter

Forex Exporter is a prometheus exporter using [OpenExchangeRates Latest Endpoint](https://docs.openexchangerates.org/docs/latest-json).
Their free tier supports 1000 calls per month.
This application has been hardcoded to hit the API once per hour which leaves ~250 requests for application restarts or testing.
It has also been hardcoded to request for USD as the base currency.

Metrics look like the following:

```
# HELP fx_rate fx rate is the exchange rate between the base and quote e.g. USD and GBP.
# TYPE fx_rate gauge
fx_rate{base="USD",quote="AED"} 3.673009
fx_rate{base="USD",quote="AFN"} 76.988
fx_rate{base="USD",quote="ALL"} 104.887255
fx_rate{base="USD",quote="AMD"} 481.616228
fx_rate{base="USD",quote="ANG"} 1.800176
fx_rate{base="USD",quote="AOA"} 582.5
fx_rate{base="USD",quote="ARS"} 72.89231
```

# Contributing

PRs welcome.

Ideas:

- CLI library
- Customisable base pair
- Customisable duration to hit the API.

