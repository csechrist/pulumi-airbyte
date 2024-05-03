// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class ConnectionConfigurationsStreamArgs : global::Pulumi.ResourceArgs
    {
        [Input("cursorFields")]
        private InputList<string>? _cursorFields;

        /// <summary>
        /// Path to the field that will be used to determine if a record is new or modified since the last sync. This field is REQUIRED if `sync_mode` is `incremental` unless there is a default.
        /// </summary>
        public InputList<string> CursorFields
        {
            get => _cursorFields ?? (_cursorFields = new InputList<string>());
            set => _cursorFields = value;
        }

        /// <summary>
        /// Not Null
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("primaryKeys")]
        private InputList<ImmutableArray<string>>? _primaryKeys;

        /// <summary>
        /// Paths to the fields that will be used as primary key. This field is REQUIRED if `destination_sync_mode` is `*_dedup` unless it is already supplied by the source schema.
        /// </summary>
        public InputList<ImmutableArray<string>> PrimaryKeys
        {
            get => _primaryKeys ?? (_primaryKeys = new InputList<ImmutableArray<string>>());
            set => _primaryKeys = value;
        }

        /// <summary>
        /// must be one of ["full*refresh*overwrite", "full*refresh*append", "incremental*append", "incremental*deduped_history"]
        /// </summary>
        [Input("syncMode")]
        public Input<string>? SyncMode { get; set; }

        public ConnectionConfigurationsStreamArgs()
        {
        }
        public static new ConnectionConfigurationsStreamArgs Empty => new ConnectionConfigurationsStreamArgs();
    }
}