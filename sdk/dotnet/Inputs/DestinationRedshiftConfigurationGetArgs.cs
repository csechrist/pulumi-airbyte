// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class DestinationRedshiftConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// Disable Writing Final Tables. WARNING! The data format in *airbyte*data is likely stable but there are no guarantees that other metadata columns will remain the same in future versions. Default: false
        /// </summary>
        [Input("disableTypeDedupe")]
        public Input<bool>? DisableTypeDedupe { get; set; }

        /// <summary>
        /// When enabled your data will load into your final tables incrementally while your data is still being synced. When Disabled (the default), your data loads into your final tables once at the end of a sync. Note that this option only applies if you elect to create Final tables. Default: false
        /// </summary>
        [Input("enableIncrementalFinalTableUpdates")]
        public Input<bool>? EnableIncrementalFinalTableUpdates { get; set; }

        /// <summary>
        /// Host Endpoint of the Redshift Cluster (must include the cluster-id, region and end with .redshift.amazonaws.com)
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&amp;'. (example: key1=value1&amp;key2=value2&amp;key3=value3).
        /// </summary>
        [Input("jdbcUrlParams")]
        public Input<string>? JdbcUrlParams { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Password associated with the username.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port of the database. Default: 5439
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The schema to write raw tables into
        /// </summary>
        [Input("rawDataSchema")]
        public Input<string>? RawDataSchema { get; set; }

        /// <summary>
        /// The default schema tables are written to if the source does not specify a namespace. Unless specifically configured, the usual value for this field is "public". Default: "public"
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
        /// </summary>
        [Input("tunnelMethod")]
        public Input<Inputs.DestinationRedshiftConfigurationTunnelMethodGetArgs>? TunnelMethod { get; set; }

        /// <summary>
        /// The way data will be uploaded to Redshift.
        /// </summary>
        [Input("uploadingMethod")]
        public Input<Inputs.DestinationRedshiftConfigurationUploadingMethodGetArgs>? UploadingMethod { get; set; }

        /// <summary>
        /// Username to use to access the database.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public DestinationRedshiftConfigurationGetArgs()
        {
        }
        public static new DestinationRedshiftConfigurationGetArgs Empty => new DestinationRedshiftConfigurationGetArgs();
    }
}
