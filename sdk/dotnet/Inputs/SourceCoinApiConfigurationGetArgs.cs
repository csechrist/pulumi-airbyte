// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceCoinApiConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        private Input<string>? _apiKey;

        /// <summary>
        /// API Key
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
        /// The end date in ISO 8601 format. If not supplied, data will be returned
        /// from the start date to the current time, or when the count of result
        /// elements reaches its limit.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// The environment to use. Either sandbox or production.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The maximum number of elements to return. If not supplied, the default
        /// is 100. For numbers larger than 100, each 100 items is counted as one
        /// request for pricing purposes. Maximum value is 100000.
        /// 
        /// Default: 100
        /// </summary>
        [Input("limit")]
        public Input<int>? Limit { get; set; }

        /// <summary>
        /// The period to use. See the documentation for a list. https://docs.coinapi.io/#list-all-periods-get
        /// </summary>
        [Input("period", required: true)]
        public Input<string> Period { get; set; } = null!;

        /// <summary>
        /// The start date in ISO 8601 format.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        /// <summary>
        /// The symbol ID to use. See the documentation for a list.
        /// https://docs.coinapi.io/#list-all-symbols-get
        /// </summary>
        [Input("symbolId", required: true)]
        public Input<string> SymbolId { get; set; } = null!;

        public SourceCoinApiConfigurationGetArgs()
        {
        }
        public static new SourceCoinApiConfigurationGetArgs Empty => new SourceCoinApiConfigurationGetArgs();
    }
}
