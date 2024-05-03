// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSmartsheetsConfigurationCredentialsApiAccessTokenGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// The access token to use for accessing your data from Smartsheets. This access token must be generated by a user with at least read access to the data you'd like to replicate. Generate an access token in the Smartsheets main menu by clicking Account &gt; Apps &amp; Integrations &gt; API Access. See the \n\nsetup guide\n\n for information on how to obtain this token.
        /// </summary>
        public Input<string>? AccessToken
        {
            get => _accessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourceSmartsheetsConfigurationCredentialsApiAccessTokenGetArgs()
        {
        }
        public static new SourceSmartsheetsConfigurationCredentialsApiAccessTokenGetArgs Empty => new SourceSmartsheetsConfigurationCredentialsApiAccessTokenGetArgs();
    }
}