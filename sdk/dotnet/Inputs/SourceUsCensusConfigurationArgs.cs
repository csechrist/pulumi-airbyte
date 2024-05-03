// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceUsCensusConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        private Input<string>? _apiKey;

        /// <summary>
        /// Your API Key. Get your key \n\nhere\n\n.
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
        /// The query parameters portion of the GET request, without the api key
        /// </summary>
        [Input("queryParams")]
        public Input<string>? QueryParams { get; set; }

        /// <summary>
        /// The path portion of the GET request
        /// </summary>
        [Input("queryPath", required: true)]
        public Input<string> QueryPath { get; set; } = null!;

        public SourceUsCensusConfigurationArgs()
        {
        }
        public static new SourceUsCensusConfigurationArgs Empty => new SourceUsCensusConfigurationArgs();
    }
}