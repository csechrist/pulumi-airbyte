// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationWeaviateConfigurationProcessingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of overlap between chunks in tokens to store in vector store to better capture relevant context. Default: 0
        /// </summary>
        [Input("chunkOverlap")]
        public Input<int>? ChunkOverlap { get; set; }

        /// <summary>
        /// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
        /// </summary>
        [Input("chunkSize", required: true)]
        public Input<int> ChunkSize { get; set; } = null!;

        [Input("fieldNameMappings")]
        private InputList<Inputs.DestinationWeaviateConfigurationProcessingFieldNameMappingGetArgs>? _fieldNameMappings;

        /// <summary>
        /// List of fields to rename. Not applicable for nested fields, but can be used to rename fields already flattened via dot notation.
        /// </summary>
        public InputList<Inputs.DestinationWeaviateConfigurationProcessingFieldNameMappingGetArgs> FieldNameMappings
        {
            get => _fieldNameMappings ?? (_fieldNameMappings = new InputList<Inputs.DestinationWeaviateConfigurationProcessingFieldNameMappingGetArgs>());
            set => _fieldNameMappings = value;
        }

        [Input("metadataFields")]
        private InputList<string>? _metadataFields;

        /// <summary>
        /// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
        /// </summary>
        public InputList<string> MetadataFields
        {
            get => _metadataFields ?? (_metadataFields = new InputList<string>());
            set => _metadataFields = value;
        }

        [Input("textFields")]
        private InputList<string>? _textFields;

        /// <summary>
        /// List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
        /// </summary>
        public InputList<string> TextFields
        {
            get => _textFields ?? (_textFields = new InputList<string>());
            set => _textFields = value;
        }

        /// <summary>
        /// Split text fields into chunks based on the specified method.
        /// </summary>
        [Input("textSplitter")]
        public Input<Inputs.DestinationWeaviateConfigurationProcessingTextSplitterGetArgs>? TextSplitter { get; set; }

        public DestinationWeaviateConfigurationProcessingGetArgs()
        {
        }
        public static new DestinationWeaviateConfigurationProcessingGetArgs Empty => new DestinationWeaviateConfigurationProcessingGetArgs();
    }
}
