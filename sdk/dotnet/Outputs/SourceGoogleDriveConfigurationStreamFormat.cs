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
    public sealed class SourceGoogleDriveConfigurationStreamFormat
    {
        public readonly Outputs.SourceGoogleDriveConfigurationStreamFormatAvroFormat? AvroFormat;
        public readonly Outputs.SourceGoogleDriveConfigurationStreamFormatCsvFormat? CsvFormat;
        /// <summary>
        /// Extract text from document formats (.pdf, .docx, .md, .pptx) and emit as one record per file.
        /// </summary>
        public readonly Outputs.SourceGoogleDriveConfigurationStreamFormatDocumentFileTypeFormatExperimental? DocumentFileTypeFormatExperimental;
        public readonly Outputs.SourceGoogleDriveConfigurationStreamFormatJsonlFormat? JsonlFormat;
        public readonly Outputs.SourceGoogleDriveConfigurationStreamFormatParquetFormat? ParquetFormat;

        [OutputConstructor]
        private SourceGoogleDriveConfigurationStreamFormat(
            Outputs.SourceGoogleDriveConfigurationStreamFormatAvroFormat? avroFormat,

            Outputs.SourceGoogleDriveConfigurationStreamFormatCsvFormat? csvFormat,

            Outputs.SourceGoogleDriveConfigurationStreamFormatDocumentFileTypeFormatExperimental? documentFileTypeFormatExperimental,

            Outputs.SourceGoogleDriveConfigurationStreamFormatJsonlFormat? jsonlFormat,

            Outputs.SourceGoogleDriveConfigurationStreamFormatParquetFormat? parquetFormat)
        {
            AvroFormat = avroFormat;
            CsvFormat = csvFormat;
            DocumentFileTypeFormatExperimental = documentFileTypeFormatExperimental;
            JsonlFormat = jsonlFormat;
            ParquetFormat = parquetFormat;
        }
    }
}