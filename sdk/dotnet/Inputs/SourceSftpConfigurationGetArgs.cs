// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceSftpConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The server authentication method
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.SourceSftpConfigurationCredentialsGetArgs>? Credentials { get; set; }

        /// <summary>
        /// The regular expression to specify files for sync in a chosen Folder Path. Default: ""
        /// </summary>
        [Input("filePattern")]
        public Input<string>? FilePattern { get; set; }

        /// <summary>
        /// Coma separated file types. Currently only 'csv' and 'json' types are supported. Default: "csv,json"
        /// </summary>
        [Input("fileTypes")]
        public Input<string>? FileTypes { get; set; }

        /// <summary>
        /// The directory to search files for sync. Default: ""
        /// </summary>
        [Input("folderPath")]
        public Input<string>? FolderPath { get; set; }

        /// <summary>
        /// The server host address
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// The server port. Default: 22
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The server user
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public SourceSftpConfigurationGetArgs()
        {
        }
        public static new SourceSftpConfigurationGetArgs Empty => new SourceSftpConfigurationGetArgs();
    }
}
