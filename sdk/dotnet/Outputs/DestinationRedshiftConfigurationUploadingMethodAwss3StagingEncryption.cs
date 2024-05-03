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
    public sealed class DestinationRedshiftConfigurationUploadingMethodAwss3StagingEncryption
    {
        /// <summary>
        /// Staging data will be encrypted using AES-CBC envelope encryption.
        /// </summary>
        public readonly Outputs.DestinationRedshiftConfigurationUploadingMethodAwss3StagingEncryptionAescbcEnvelopeEncryption? AescbcEnvelopeEncryption;
        /// <summary>
        /// Staging data will be stored in plaintext.
        /// </summary>
        public readonly Outputs.DestinationRedshiftConfigurationUploadingMethodAwss3StagingEncryptionNoEncryption? NoEncryption;

        [OutputConstructor]
        private DestinationRedshiftConfigurationUploadingMethodAwss3StagingEncryption(
            Outputs.DestinationRedshiftConfigurationUploadingMethodAwss3StagingEncryptionAescbcEnvelopeEncryption? aescbcEnvelopeEncryption,

            Outputs.DestinationRedshiftConfigurationUploadingMethodAwss3StagingEncryptionNoEncryption? noEncryption)
        {
            AescbcEnvelopeEncryption = aescbcEnvelopeEncryption;
            NoEncryption = noEncryption;
        }
    }
}