// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationXataConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey", required: true)]
        private Input<string>? _apiKey;

        /// <summary>
        /// API Key to connect.
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
        /// URL pointing to your workspace.
        /// </summary>
        [Input("dbUrl", required: true)]
        public Input<string> DbUrl { get; set; } = null!;

        public DestinationXataConfigurationGetArgs()
        {
        }
        public static new DestinationXataConfigurationGetArgs Empty => new DestinationXataConfigurationGetArgs();
    }
}