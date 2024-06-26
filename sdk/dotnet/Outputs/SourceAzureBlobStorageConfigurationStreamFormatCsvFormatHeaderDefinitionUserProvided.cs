// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Outputs
{

    [OutputType]
    public sealed class SourceAzureBlobStorageConfigurationStreamFormatCsvFormatHeaderDefinitionUserProvided
    {
        /// <summary>
        /// The column names that will be used while emitting the CSV records
        /// </summary>
        public readonly ImmutableArray<string> ColumnNames;

        [OutputConstructor]
        private SourceAzureBlobStorageConfigurationStreamFormatCsvFormatHeaderDefinitionUserProvided(ImmutableArray<string> columnNames)
        {
            ColumnNames = columnNames;
        }
    }
}
