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
    public sealed class SourceZendeskChatConfiguration
    {
        public readonly Outputs.SourceZendeskChatConfigurationCredentials? Credentials;
        /// <summary>
        /// The date from which you'd like to replicate data for Zendesk Chat API, in the format YYYY-MM-DDT00:00:00Z.
        /// </summary>
        public readonly string StartDate;
        /// <summary>
        /// Required if you access Zendesk Chat from a Zendesk Support subdomain. Default: ""
        /// </summary>
        public readonly string? Subdomain;

        [OutputConstructor]
        private SourceZendeskChatConfiguration(
            Outputs.SourceZendeskChatConfigurationCredentials? credentials,

            string startDate,

            string? subdomain)
        {
            Credentials = credentials;
            StartDate = startDate;
            Subdomain = subdomain;
        }
    }
}