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
    public sealed class SourceLinkedinPagesConfigurationCredentialsAccessToken
    {
        /// <summary>
        /// The token value generated using the LinkedIn Developers OAuth Token Tools. See the \n\ndocs\n\n to obtain yours.
        /// </summary>
        public readonly string AccessToken;

        [OutputConstructor]
        private SourceLinkedinPagesConfigurationCredentialsAccessToken(string accessToken)
        {
            AccessToken = accessToken;
        }
    }
}