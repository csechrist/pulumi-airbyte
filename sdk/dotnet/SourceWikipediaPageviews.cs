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
    /// SourceWikipediaPageviews Resource
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
    ///     var mySourceWikipediapageviews = new Airbyte.SourceWikipediaPageviews("mySourceWikipediapageviews", new()
    ///     {
    ///         Configuration = new Airbyte.Inputs.SourceWikipediaPageviewsConfigurationArgs
    ///         {
    ///             Access = "all-access",
    ///             Agent = "user",
    ///             Article = "Are_You_the_One%3F",
    ///             Country = "IN",
    ///             End = "...my_end...",
    ///             Project = "www.mediawiki.org",
    ///             Start = "...my_start...",
    ///         },
    ///         DefinitionId = "54523f36-dab5-4122-890f-3e992c2a3f4c",
    ///         SecretId = "...my_secret_id...",
    ///         WorkspaceId = "c7cc4eaf-dab4-4c1b-8af6-6c12869f984d",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AirbyteResourceType("airbyte:index/sourceWikipediaPageviews:SourceWikipediaPageviews")]
    public partial class SourceWikipediaPageviews : global::Pulumi.CustomResource
    {
        [Output("configuration")]
        public Output<Outputs.SourceWikipediaPageviewsConfiguration> Configuration { get; private set; } = null!;

        /// <summary>
        /// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
        /// </summary>
        [Output("definitionId")]
        public Output<string?> DefinitionId { get; private set; } = null!;

        /// <summary>
        /// Name of the source e.g. dev-mysql-instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
        /// </summary>
        [Output("secretId")]
        public Output<string?> SecretId { get; private set; } = null!;

        [Output("sourceId")]
        public Output<string> SourceId { get; private set; } = null!;

        [Output("sourceType")]
        public Output<string> SourceType { get; private set; } = null!;

        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a SourceWikipediaPageviews resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SourceWikipediaPageviews(string name, SourceWikipediaPageviewsArgs args, CustomResourceOptions? options = null)
            : base("airbyte:index/sourceWikipediaPageviews:SourceWikipediaPageviews", name, args ?? new SourceWikipediaPageviewsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SourceWikipediaPageviews(string name, Input<string> id, SourceWikipediaPageviewsState? state = null, CustomResourceOptions? options = null)
            : base("airbyte:index/sourceWikipediaPageviews:SourceWikipediaPageviews", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SourceWikipediaPageviews resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SourceWikipediaPageviews Get(string name, Input<string> id, SourceWikipediaPageviewsState? state = null, CustomResourceOptions? options = null)
        {
            return new SourceWikipediaPageviews(name, id, state, options);
        }
    }

    public sealed class SourceWikipediaPageviewsArgs : global::Pulumi.ResourceArgs
    {
        [Input("configuration", required: true)]
        public Input<Inputs.SourceWikipediaPageviewsConfigurationArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
        /// </summary>
        [Input("definitionId")]
        public Input<string>? DefinitionId { get; set; }

        /// <summary>
        /// Name of the source e.g. dev-mysql-instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public SourceWikipediaPageviewsArgs()
        {
        }
        public static new SourceWikipediaPageviewsArgs Empty => new SourceWikipediaPageviewsArgs();
    }

    public sealed class SourceWikipediaPageviewsState : global::Pulumi.ResourceArgs
    {
        [Input("configuration")]
        public Input<Inputs.SourceWikipediaPageviewsConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
        /// </summary>
        [Input("definitionId")]
        public Input<string>? DefinitionId { get; set; }

        /// <summary>
        /// Name of the source e.g. dev-mysql-instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public SourceWikipediaPageviewsState()
        {
        }
        public static new SourceWikipediaPageviewsState Empty => new SourceWikipediaPageviewsState();
    }
}
