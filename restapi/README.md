# REST-accessible oracle

This example oracle is REST-accessible. It automatically generates R-points on request and keeps published values stored for later retrieval. 

The data sources publish every 5 minutes.

This project can serve as an oracle while forming Discreet Log Contracts. A live version of this oracle is running on [https://oracle.gertjaap.org/] 

## REST Endpoints

| resource          | description                              |
|:------------------|:-----------------------------------------|
|[`/api/pubkey`](https://oracle.gertjaap.org/api/pubkey)      | Returns the public keys of the oracle     |
|[`/api/datasources`](https://oracle.gertjaap.org/api/datasources) | Returns an array of data sources the oracle publishes |
|[`/api/rpoint/{s}/{t}`](https://oracle.gertjaap.org/api/rpoint/1/1527494400) | Returns the public one-time-signing key for datasource with ID **s** at the unix timestamp **t**. |
|[`/api/publication/{R}`](https://oracle.gertjaap.org/api/publication/03bb906ac1e3926fdbb305d710eb7878720b70f628e1e0141a21be36101c93982e) | Returns the value and signature published for data source point **R** (if published). R is hex encoded [33]byte |

## Using the public deployment

You're free to use my public deployment of the oracle as well. I have linked the URLs of the public deployment in the REST endpoint table above.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
