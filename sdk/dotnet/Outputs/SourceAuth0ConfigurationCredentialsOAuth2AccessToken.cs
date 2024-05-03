// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class SourceAuth0ConfigurationCredentialsOAuth2AccessToken
    {
        /// <summary>
        /// Also called \n\nAPI Access Token \n\n The access token used to call the Auth0 Management API Token. It's a JWT that contains specific grant permissions knowns as scopes.
        /// </summary>
        public readonly string AccessToken;

        [OutputConstructor]
        private SourceAuth0ConfigurationCredentialsOAuth2AccessToken(string accessToken)
        {
            AccessToken = accessToken;
        }
    }
}