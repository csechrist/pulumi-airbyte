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
    public sealed class SourceHarvestConfiguration
    {
        /// <summary>
        /// Harvest account ID. Required for all Harvest requests in pair with Personal Access Token
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// Choose how to authenticate to Harvest.
        /// </summary>
        public readonly Outputs.SourceHarvestConfigurationCredentials? Credentials;
        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00Z. Any data after this date will not be replicated.
        /// </summary>
        public readonly string? ReplicationEndDate;
        /// <summary>
        /// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
        /// </summary>
        public readonly string ReplicationStartDate;

        [OutputConstructor]
        private SourceHarvestConfiguration(
            string accountId,

            Outputs.SourceHarvestConfigurationCredentials? credentials,

            string? replicationEndDate,

            string replicationStartDate)
        {
            AccountId = accountId;
            Credentials = credentials;
            ReplicationEndDate = replicationEndDate;
            ReplicationStartDate = replicationStartDate;
        }
    }
}
