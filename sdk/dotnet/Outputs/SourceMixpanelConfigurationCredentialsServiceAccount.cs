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
    public sealed class SourceMixpanelConfigurationCredentialsServiceAccount
    {
        /// <summary>
        /// Your project ID number. See the \n\ndocs\n\n for more information on how to obtain this.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// Mixpanel Service Account Secret. See the \n\ndocs\n\n for more information on how to obtain this.
        /// </summary>
        public readonly string Secret;
        /// <summary>
        /// Mixpanel Service Account Username. See the \n\ndocs\n\n for more information on how to obtain this.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private SourceMixpanelConfigurationCredentialsServiceAccount(
            int projectId,

            string secret,

            string username)
        {
            ProjectId = projectId;
            Secret = secret;
            Username = username;
        }
    }
}
