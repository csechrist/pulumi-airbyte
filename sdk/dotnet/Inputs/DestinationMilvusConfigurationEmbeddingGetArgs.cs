// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationMilvusConfigurationEmbeddingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use the Azure-hosted OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
        /// </summary>
        [Input("azureOpenAi")]
        public Input<Inputs.DestinationMilvusConfigurationEmbeddingAzureOpenAiGetArgs>? AzureOpenAi { get; set; }

        /// <summary>
        /// Use the Cohere API to embed text.
        /// </summary>
        [Input("cohere")]
        public Input<Inputs.DestinationMilvusConfigurationEmbeddingCohereGetArgs>? Cohere { get; set; }

        /// <summary>
        /// Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
        /// </summary>
        [Input("fake")]
        public Input<Inputs.DestinationMilvusConfigurationEmbeddingFakeGetArgs>? Fake { get; set; }

        /// <summary>
        /// Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
        /// </summary>
        [Input("openAi")]
        public Input<Inputs.DestinationMilvusConfigurationEmbeddingOpenAiGetArgs>? OpenAi { get; set; }

        /// <summary>
        /// Use a service that's compatible with the OpenAI API to embed text.
        /// </summary>
        [Input("openAiCompatible")]
        public Input<Inputs.DestinationMilvusConfigurationEmbeddingOpenAiCompatibleGetArgs>? OpenAiCompatible { get; set; }

        public DestinationMilvusConfigurationEmbeddingGetArgs()
        {
        }
        public static new DestinationMilvusConfigurationEmbeddingGetArgs Empty => new DestinationMilvusConfigurationEmbeddingGetArgs();
    }
}