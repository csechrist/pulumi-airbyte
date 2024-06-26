// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceS3ConfigurationStreamFormatArgs : global::Pulumi.ResourceArgs
    {
        [Input("avroFormat")]
        public Input<Inputs.SourceS3ConfigurationStreamFormatAvroFormatArgs>? AvroFormat { get; set; }

        [Input("csvFormat")]
        public Input<Inputs.SourceS3ConfigurationStreamFormatCsvFormatArgs>? CsvFormat { get; set; }

        /// <summary>
        /// Extract text from document formats (.pdf, .docx, .md, .pptx) and emit as one record per file.
        /// </summary>
        [Input("documentFileTypeFormatExperimental")]
        public Input<Inputs.SourceS3ConfigurationStreamFormatDocumentFileTypeFormatExperimentalArgs>? DocumentFileTypeFormatExperimental { get; set; }

        [Input("jsonlFormat")]
        public Input<Inputs.SourceS3ConfigurationStreamFormatJsonlFormatArgs>? JsonlFormat { get; set; }

        [Input("parquetFormat")]
        public Input<Inputs.SourceS3ConfigurationStreamFormatParquetFormatArgs>? ParquetFormat { get; set; }

        public SourceS3ConfigurationStreamFormatArgs()
        {
        }
        public static new SourceS3ConfigurationStreamFormatArgs Empty => new SourceS3ConfigurationStreamFormatArgs();
    }
}
