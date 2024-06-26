// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMailgunConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain region code. 'EU' or 'US' are possible values. The default is 'US'. Default: "US"
        /// </summary>
        [Input("domainRegion")]
        public Input<string>? DomainRegion { get; set; }

        [Input("privateKey", required: true)]
        private Input<string>? _privateKey;

        /// <summary>
        /// Primary account API key to access your Mailgun data.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// UTC date and time in the format 2020-10-01 00:00:00. Any data before this date will not be replicated. If omitted, defaults to 3 days ago.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        public SourceMailgunConfigurationGetArgs()
        {
        }
        public static new SourceMailgunConfigurationGetArgs Empty => new SourceMailgunConfigurationGetArgs();
    }
}
