// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSlackConfigurationCredentialsSignInViaSlackOAuthGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessToken", required: true)]
        private Input<string>? _accessToken;

        /// <summary>
        /// Slack access_token. See our \n\ndocs\n\n if you need help generating the token.
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
        /// Slack client_id. See our \n\ndocs\n\n if you need help finding this id.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Slack client_secret. See our \n\ndocs\n\n if you need help finding this secret.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        public SourceSlackConfigurationCredentialsSignInViaSlackOAuthGetArgs()
        {
        }
        public static new SourceSlackConfigurationCredentialsSignInViaSlackOAuthGetArgs Empty => new SourceSlackConfigurationCredentialsSignInViaSlackOAuthGetArgs();
    }
}
