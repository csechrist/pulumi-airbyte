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
    /// DestinationWeaviate Resource
    /// </summary>
    [AirbyteResourceType("airbyte:index/destinationWeaviate:DestinationWeaviate")]
    public partial class DestinationWeaviate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination configuration,
        /// as well as to provide type safety for the configuration passed to the destination.
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.DestinationWeaviateConfiguration> Configuration { get; private set; } = null!;

        /// <summary>
        /// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
        /// replacement if changed.
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
        /// Create a DestinationWeaviate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DestinationWeaviate(string name, DestinationWeaviateArgs args, CustomResourceOptions? options = null)
            : base("airbyte:index/destinationWeaviate:DestinationWeaviate", name, args ?? new DestinationWeaviateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DestinationWeaviate(string name, Input<string> id, DestinationWeaviateState? state = null, CustomResourceOptions? options = null)
            : base("airbyte:index/destinationWeaviate:DestinationWeaviate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DestinationWeaviate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DestinationWeaviate Get(string name, Input<string> id, DestinationWeaviateState? state = null, CustomResourceOptions? options = null)
        {
            return new DestinationWeaviate(name, id, state, options);
        }
    }

    public sealed class DestinationWeaviateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination configuration,
        /// as well as to provide type safety for the configuration passed to the destination.
        /// </summary>
        [Input("configuration", required: true)]
        public Input<Inputs.DestinationWeaviateConfigurationArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
        /// replacement if changed.
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

        public DestinationWeaviateArgs()
        {
        }
        public static new DestinationWeaviateArgs Empty => new DestinationWeaviateArgs();
    }

    public sealed class DestinationWeaviateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination configuration,
        /// as well as to provide type safety for the configuration passed to the destination.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.DestinationWeaviateConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires
        /// replacement if changed.
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

        public DestinationWeaviateState()
        {
        }
        public static new DestinationWeaviateState Empty => new DestinationWeaviateState();
    }
}
