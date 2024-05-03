// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceCoinmarketcapConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        private Input<string>? _apiKey;

        /// <summary>
        /// Your API Key. See \n\nhere\n\n. The token is case sensitive.
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// /latest: Latest market ticker quotes and averages for cryptocurrencies and exchanges. /historical: Intervals of historic market data like OHLCV data or data for use in charting libraries. See \n\nhere\n\n. must be one of ["latest", "historical"]
        /// </summary>
        [Input("dataType", required: true)]
        public Input<string> DataType { get; set; } = null!;

        [Input("symbols")]
        private InputList<string>? _symbols;

        /// <summary>
        /// Cryptocurrency symbols. (only used for quotes stream)
        /// </summary>
        public InputList<string> Symbols
        {
            get => _symbols ?? (_symbols = new InputList<string>());
            set => _symbols = value;
        }

        public SourceCoinmarketcapConfigurationGetArgs()
        {
        }
        public static new SourceCoinmarketcapConfigurationGetArgs Empty => new SourceCoinmarketcapConfigurationGetArgs();
    }
}