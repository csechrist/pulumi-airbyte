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
    public sealed class DestinationGcsConfigurationCredentialHmacKey
    {
        /// <summary>
        /// must be one of ["HMAC*KEY"]; Default: "HMAC*KEY"
        /// </summary>
        public readonly string? CredentialType;
        /// <summary>
        /// When linked to a service account, this ID is 61 characters long; when linked to a user account, it is 24 characters long. Read more \n\nhere\n\n.
        /// </summary>
        public readonly string HmacKeyAccessId;
        /// <summary>
        /// The corresponding secret for the access ID. It is a 40-character base-64 encoded string.  Read more \n\nhere\n\n.
        /// </summary>
        public readonly string HmacKeySecret;

        [OutputConstructor]
        private DestinationGcsConfigurationCredentialHmacKey(
            string? credentialType,

            string hmacKeyAccessId,

            string hmacKeySecret)
        {
            CredentialType = credentialType;
            HmacKeyAccessId = hmacKeyAccessId;
            HmacKeySecret = hmacKeySecret;
        }
    }
}