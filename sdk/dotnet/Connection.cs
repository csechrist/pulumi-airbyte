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
    /// Connection Resource
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
    ///     var myConnection = new Airbyte.Connection("myConnection", new()
    ///     {
    ///         DataResidency = "eu",
    ///         DestinationId = "669dd1e3-6208-43ea-bc85-5914e0a570f6",
    ///         NamespaceDefinition = "custom_format",
    ///         NamespaceFormat = SOURCE_NAMESPACE,
    ///         NonBreakingSchemaUpdatesBehavior = "propagate_columns",
    ///         Prefix = "...my_prefix...",
    ///         SourceId = "3a555847-8358-4423-a5b6-c7b3fd2fd307",
    ///         Status = "deprecated",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [AirbyteResourceType("airbyte:index/connection:Connection")]
    public partial class Connection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of configured stream options for a connection.
        /// </summary>
        [Output("configurations")]
        public Output<Outputs.ConnectionConfigurations> Configurations { get; private set; } = null!;

        [Output("connectionId")]
        public Output<string> ConnectionId { get; private set; } = null!;

        /// <summary>
        /// must be one of ["auto", "us", "eu"]; Default: "auto"
        /// </summary>
        [Output("dataResidency")]
        public Output<string> DataResidency { get; private set; } = null!;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Output("destinationId")]
        public Output<string> DestinationId { get; private set; } = null!;

        /// <summary>
        /// Optional name of the connection
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Define the location where the data will be stored in the destination. must be one of ["source", "destination", "custom_format"]; Default: "destination"
        /// </summary>
        [Output("namespaceDefinition")]
        public Output<string> NamespaceDefinition { get; private set; } = null!;

        /// <summary>
        /// Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'. Default: null
        /// </summary>
        [Output("namespaceFormat")]
        public Output<string> NamespaceFormat { get; private set; } = null!;

        /// <summary>
        /// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source. must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]; Default: "ignore"
        /// </summary>
        [Output("nonBreakingSchemaUpdatesBehavior")]
        public Output<string> NonBreakingSchemaUpdatesBehavior { get; private set; } = null!;

        /// <summary>
        /// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” =&gt; “airbyte*projects”).
        /// </summary>
        [Output("prefix")]
        public Output<string> Prefix { get; private set; } = null!;

        /// <summary>
        /// schedule for when the the connection should run, per the schedule type
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.ConnectionSchedule> Schedule { get; private set; } = null!;

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Output("sourceId")]
        public Output<string> SourceId { get; private set; } = null!;

        /// <summary>
        /// must be one of ["active", "inactive", "deprecated"]
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a Connection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connection(string name, ConnectionArgs args, CustomResourceOptions? options = null)
            : base("airbyte:index/connection:Connection", name, args ?? new ConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Connection(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
            : base("airbyte:index/connection:Connection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Connection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connection Get(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new Connection(name, id, state, options);
        }
    }

    public sealed class ConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A list of configured stream options for a connection.
        /// </summary>
        [Input("configurations")]
        public Input<Inputs.ConnectionConfigurationsArgs>? Configurations { get; set; }

        /// <summary>
        /// must be one of ["auto", "us", "eu"]; Default: "auto"
        /// </summary>
        [Input("dataResidency")]
        public Input<string>? DataResidency { get; set; }

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Input("destinationId", required: true)]
        public Input<string> DestinationId { get; set; } = null!;

        /// <summary>
        /// Optional name of the connection
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Define the location where the data will be stored in the destination. must be one of ["source", "destination", "custom_format"]; Default: "destination"
        /// </summary>
        [Input("namespaceDefinition")]
        public Input<string>? NamespaceDefinition { get; set; }

        /// <summary>
        /// Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'. Default: null
        /// </summary>
        [Input("namespaceFormat")]
        public Input<string>? NamespaceFormat { get; set; }

        /// <summary>
        /// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source. must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]; Default: "ignore"
        /// </summary>
        [Input("nonBreakingSchemaUpdatesBehavior")]
        public Input<string>? NonBreakingSchemaUpdatesBehavior { get; set; }

        /// <summary>
        /// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” =&gt; “airbyte*projects”).
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// schedule for when the the connection should run, per the schedule type
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ConnectionScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        /// <summary>
        /// must be one of ["active", "inactive", "deprecated"]
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ConnectionArgs()
        {
        }
        public static new ConnectionArgs Empty => new ConnectionArgs();
    }

    public sealed class ConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A list of configured stream options for a connection.
        /// </summary>
        [Input("configurations")]
        public Input<Inputs.ConnectionConfigurationsGetArgs>? Configurations { get; set; }

        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        /// <summary>
        /// must be one of ["auto", "us", "eu"]; Default: "auto"
        /// </summary>
        [Input("dataResidency")]
        public Input<string>? DataResidency { get; set; }

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Input("destinationId")]
        public Input<string>? DestinationId { get; set; }

        /// <summary>
        /// Optional name of the connection
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Define the location where the data will be stored in the destination. must be one of ["source", "destination", "custom_format"]; Default: "destination"
        /// </summary>
        [Input("namespaceDefinition")]
        public Input<string>? NamespaceDefinition { get; set; }

        /// <summary>
        /// Used when namespaceDefinition is 'custom*format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE*NAMESPACE}" then behaves like namespaceDefinition = 'source'. Default: null
        /// </summary>
        [Input("namespaceFormat")]
        public Input<string>? NamespaceFormat { get; set; }

        /// <summary>
        /// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source. must be one of ["ignore", "disable*connection", "propagate*columns", "propagate_fully"]; Default: "ignore"
        /// </summary>
        [Input("nonBreakingSchemaUpdatesBehavior")]
        public Input<string>? NonBreakingSchemaUpdatesBehavior { get; set; }

        /// <summary>
        /// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte*” causes “projects” =&gt; “airbyte*projects”).
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// schedule for when the the connection should run, per the schedule type
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ConnectionScheduleGetArgs>? Schedule { get; set; }

        /// <summary>
        /// Requires replacement if changed.
        /// </summary>
        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        /// <summary>
        /// must be one of ["active", "inactive", "deprecated"]
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public ConnectionState()
        {
        }
        public static new ConnectionState Empty => new ConnectionState();
    }
}
