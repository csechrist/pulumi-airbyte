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
    /// DestinationDatabend Resource
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
    ///     var myDestinationDatabend = new Airbyte.DestinationDatabend("myDestinationDatabend", new()
    ///     {
    ///         Configuration = new Airbyte.Inputs.DestinationDatabendConfigurationArgs
    ///         {
    ///             Database = "...my_database...",
    ///             Host = "...my_host...",
    ///             Password = "...my_password...",
    ///             Port = 443,
    ///             Table = "default",
    ///             Username = "Anjali_McLaughlin1",
    ///         },
    ///         DefinitionId = "2e07129d-4644-4f9d-93d5-4c7cfb82ef1e",
    ///         WorkspaceId = "7477c9e2-c85c-4904-a203-ff157a47112d",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AirbyteResourceType("airbyte:index/destinationDatabend:DestinationDatabend")]
    public partial class DestinationDatabend : global::Pulumi.CustomResource
    {
        [Output("configuration")]
        public Output<Outputs.DestinationDatabendConfiguration> Configuration { get; private set; } = null!;

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
        /// Create a DestinationDatabend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DestinationDatabend(string name, DestinationDatabendArgs args, CustomResourceOptions? options = null)
            : base("airbyte:index/destinationDatabend:DestinationDatabend", name, args ?? new DestinationDatabendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DestinationDatabend(string name, Input<string> id, DestinationDatabendState? state = null, CustomResourceOptions? options = null)
            : base("airbyte:index/destinationDatabend:DestinationDatabend", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DestinationDatabend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DestinationDatabend Get(string name, Input<string> id, DestinationDatabendState? state = null, CustomResourceOptions? options = null)
        {
            return new DestinationDatabend(name, id, state, options);
        }
    }

    public sealed class DestinationDatabendArgs : global::Pulumi.ResourceArgs
    {
        [Input("configuration", required: true)]
        public Input<Inputs.DestinationDatabendConfigurationArgs> Configuration { get; set; } = null!;

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

        public DestinationDatabendArgs()
        {
        }
        public static new DestinationDatabendArgs Empty => new DestinationDatabendArgs();
    }

    public sealed class DestinationDatabendState : global::Pulumi.ResourceArgs
    {
        [Input("configuration")]
        public Input<Inputs.DestinationDatabendConfigurationGetArgs>? Configuration { get; set; }

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

        public DestinationDatabendState()
        {
        }
        public static new DestinationDatabendState Empty => new DestinationDatabendState();
    }
}
