// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceHubspotConfigurationCredentialsOAuthGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Client ID of your HubSpot developer application. See the \n\nHubspot docs\n\n if you need help finding this ID.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The client secret for your HubSpot developer application. See the \n\nHubspot docs\n\n if you need help finding this secret.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        [Input("refreshToken", required: true)]
        private Input<string>? _refreshToken;

        /// <summary>
        /// Refresh token to renew an expired access token. See the \n\nHubspot docs\n\n if you need help finding this token.
        /// </summary>
        public Input<string>? RefreshToken
        {
            get => _refreshToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _refreshToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourceHubspotConfigurationCredentialsOAuthGetArgs()
        {
        }
        public static new SourceHubspotConfigurationCredentialsOAuthGetArgs Empty => new SourceHubspotConfigurationCredentialsOAuthGetArgs();
    }
}