// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte
{
    /// <summary>
    /// DestinationLangchain Resource
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Airbyte = Pulumi.Airbyte;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myDestinationLangchain = new Airbyte.DestinationLangchain("myDestinationLangchain", new()
    ///     {
    ///         Configuration = new Airbyte.Inputs.DestinationLangchainConfigurationArgs
    ///         {
    ///             Embedding = new Airbyte.Inputs.DestinationLangchainConfigurationEmbeddingArgs
    ///             {
    ///                 Fake = null,
    ///             },
    ///             Indexing = new Airbyte.Inputs.DestinationLangchainConfigurationIndexingArgs
    ///             {
    ///                 ChromaLocalPersistance = new Airbyte.Inputs.DestinationLangchainConfigurationIndexingChromaLocalPersistanceArgs
    ///                 {
    ///                     CollectionName = "...my_collection_name...",
    ///                     DestinationPath = "/local/my_chroma_db",
    ///                 },
    ///             },
    ///             Processing = new Airbyte.Inputs.DestinationLangchainConfigurationProcessingArgs
    ///             {
    ///                 ChunkOverlap = 1,
    ///                 ChunkSize = 5,
    ///                 TextFields = new[]
    ///                 {
    ///                     "...",
    ///                 },
    ///             },
    ///         },
    ///         DefinitionId = "2d15ef4e-895c-4921-a618-452d1432f338",
    ///         WorkspaceId = "4ca4c8c4-bf88-4272-9c3c-6bc39a6d3f39",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AirbyteResourceType("airbyte:index/destinationLangchain:DestinationLangchain")]
    public partial class DestinationLangchain : global::Pulumi.CustomResource
    {
        [Output("configuration")]
        public Output<Outputs.DestinationLangchainConfiguration> Configuration { get; private set; } = null!;

        /// <summary>
        /// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.
        /// </summary>
        [Output("definitionId")]
        public Output<string?> DefinitionId { get; private set; } = null!;

        [Output("destinationId")]
        public Output<string> DestinationId { get; private set; } = null!;

        [Output("destinationType")]
        public Output<string> DestinationType { get; private set; } = null!;

        /// <summary>
        /// Name of the destination e.g. dev-mysql-instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a DestinationLangchain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DestinationLangchain(string name, DestinationLangchainArgs args, CustomResourceOptions? options = null)
            : base("airbyte:index/destinationLangchain:DestinationLangchain", name, args ?? new DestinationLangchainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DestinationLangchain(string name, Input<string> id, DestinationLangchainState? state = null, CustomResourceOptions? options = null)
            : base("airbyte:index/destinationLangchain:DestinationLangchain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/csechrist/pulumi-airbyte",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DestinationLangchain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DestinationLangchain Get(string name, Input<string> id, DestinationLangchainState? state = null, CustomResourceOptions? options = null)
        {
            return new DestinationLangchain(name, id, state, options);
        }
    }

    public sealed class DestinationLangchainArgs : global::Pulumi.ResourceArgs
    {
        [Input("configuration", required: true)]
        public Input<Inputs.DestinationLangchainConfigurationArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.
        /// </summary>
        [Input("definitionId")]
        public Input<string>? DefinitionId { get; set; }

        /// <summary>
        /// Name of the destination e.g. dev-mysql-instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public DestinationLangchainArgs()
        {
        }
        public static new DestinationLangchainArgs Empty => new DestinationLangchainArgs();
    }

    public sealed class DestinationLangchainState : global::Pulumi.ResourceArgs
    {
        [Input("configuration")]
        public Input<Inputs.DestinationLangchainConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed.
        /// </summary>
        [Input("definitionId")]
        public Input<string>? DefinitionId { get; set; }

        [Input("destinationId")]
        public Input<string>? DestinationId { get; set; }

        [Input("destinationType")]
        public Input<string>? DestinationType { get; set; }

        /// <summary>
        /// Name of the destination e.g. dev-mysql-instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public DestinationLangchainState()
        {
        }
        public static new DestinationLangchainState Empty => new DestinationLangchainState();
    }
}
