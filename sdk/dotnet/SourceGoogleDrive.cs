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
    /// SourceGoogleDrive Resource
    /// </summary>
    [AirbyteResourceType("airbyte:index/sourceGoogleDrive:SourceGoogleDrive")]
    public partial class SourceGoogleDrive : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Used during spec; allows the developer to configure the cloud provider specific options
        /// that are needed when users configure a file-based source.
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.SourceGoogleDriveConfiguration> Configuration { get; private set; } = null!;

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
        /// Create a SourceGoogleDrive resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SourceGoogleDrive(string name, SourceGoogleDriveArgs args, CustomResourceOptions? options = null)
            : base("airbyte:index/sourceGoogleDrive:SourceGoogleDrive", name, args ?? new SourceGoogleDriveArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SourceGoogleDrive(string name, Input<string> id, SourceGoogleDriveState? state = null, CustomResourceOptions? options = null)
            : base("airbyte:index/sourceGoogleDrive:SourceGoogleDrive", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SourceGoogleDrive resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SourceGoogleDrive Get(string name, Input<string> id, SourceGoogleDriveState? state = null, CustomResourceOptions? options = null)
        {
            return new SourceGoogleDrive(name, id, state, options);
        }
    }

    public sealed class SourceGoogleDriveArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used during spec; allows the developer to configure the cloud provider specific options
        /// that are needed when users configure a file-based source.
        /// </summary>
        [Input("configuration", required: true)]
        public Input<Inputs.SourceGoogleDriveConfigurationArgs> Configuration { get; set; } = null!;

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

        public SourceGoogleDriveArgs()
        {
        }
        public static new SourceGoogleDriveArgs Empty => new SourceGoogleDriveArgs();
    }

    public sealed class SourceGoogleDriveState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used during spec; allows the developer to configure the cloud provider specific options
        /// that are needed when users configure a file-based source.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.SourceGoogleDriveConfigurationGetArgs>? Configuration { get; set; }

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

        public SourceGoogleDriveState()
        {
        }
        public static new SourceGoogleDriveState Empty => new SourceGoogleDriveState();
    }
}
