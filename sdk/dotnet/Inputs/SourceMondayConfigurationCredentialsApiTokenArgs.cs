// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMondayConfigurationCredentialsApiTokenArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiToken", required: true)]
        private Input<string>? _apiToken;

        /// <summary>
        /// API Token for making authenticated requests.
        /// </summary>
        public Input<string>? ApiToken
        {
            get => _apiToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourceMondayConfigurationCredentialsApiTokenArgs()
        {
        }
        public static new SourceMondayConfigurationCredentialsApiTokenArgs Empty => new SourceMondayConfigurationCredentialsApiTokenArgs();
    }
}
