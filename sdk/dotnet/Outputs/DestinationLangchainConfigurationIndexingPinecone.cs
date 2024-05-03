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
    public sealed class DestinationLangchainConfigurationIndexingPinecone
    {
        /// <summary>
        /// Pinecone index to use
        /// </summary>
        public readonly string Index;
        /// <summary>
        /// Pinecone environment to use
        /// </summary>
        public readonly string PineconeEnvironment;
        public readonly string PineconeKey;

        [OutputConstructor]
        private DestinationLangchainConfigurationIndexingPinecone(
            string index,

            string pineconeEnvironment,

            string pineconeKey)
        {
            Index = index;
            PineconeEnvironment = pineconeEnvironment;
            PineconeKey = pineconeKey;
        }
    }
}