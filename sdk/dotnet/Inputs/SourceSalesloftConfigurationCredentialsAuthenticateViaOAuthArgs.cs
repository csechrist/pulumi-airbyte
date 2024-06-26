// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSalesloftConfigurationCredentialsAuthenticateViaOAuthArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// Access Token for making authenticated requests.
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

        /// <summary>
        /// The Client ID of your Salesloft developer application.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The Client Secret of your Salesloft developer application.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        [Input("refreshToken", required: true)]
        private Input<string>? _refreshToken;

        /// <summary>
        /// The token for obtaining a new access token.
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

        [Input("tokenExpiryDate", required: true)]
        private Input<string>? _tokenExpiryDate;

        /// <summary>
        /// The date-time when the access token should be refreshed.
        /// </summary>
        public Input<string>? TokenExpiryDate
        {
            get => _tokenExpiryDate;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _tokenExpiryDate = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SourceSalesloftConfigurationCredentialsAuthenticateViaOAuthArgs()
        {
        }
        public static new SourceSalesloftConfigurationCredentialsAuthenticateViaOAuthArgs Empty => new SourceSalesloftConfigurationCredentialsAuthenticateViaOAuthArgs();
    }
}
