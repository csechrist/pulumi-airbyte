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
    public sealed class SourceConfluenceConfiguration
    {
        /// <summary>
        /// Please follow the Jira confluence for generating an API token: \n\ngenerating an API token\n\n.
        /// </summary>
        public readonly string ApiToken;
        /// <summary>
        /// Your Confluence domain name
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// Your Confluence login email
        /// </summary>
        public readonly string Email;

        [OutputConstructor]
        private SourceConfluenceConfiguration(
            string apiToken,

            string domainName,

            string email)
        {
            ApiToken = apiToken;
            DomainName = domainName;
            Email = email;
        }
    }
}