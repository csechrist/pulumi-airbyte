// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceMicrosoftSharepointConfigurationStreamFormatDocumentFileTypeFormatExperimentalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Processing configuration
        /// </summary>
        [Input("processing")]
        public Input<Inputs.SourceMicrosoftSharepointConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingArgs>? Processing { get; set; }

        /// <summary>
        /// If true, skip files that cannot be parsed and pass the error message along as the _ab_source_file_parse_error field. If false, fail the sync. Default: true
        /// </summary>
        [Input("skipUnprocessableFiles")]
        public Input<bool>? SkipUnprocessableFiles { get; set; }

        /// <summary>
        /// The strategy used to parse documents. `fast` extracts text directly from the document which doesn't work for all files. `ocr_only` is more reliable, but slower. `hi_res` is the most reliable, but requires an API key and a hosted instance of unstructured and can't be used with local mode. See the unstructured.io documentation for more details: https://unstructured-io.github.io/unstructured/core/partition.html#partition-pdf. must be one of ["auto", "fast", "ocr_only", "hi_res"]; Default: "auto"
        /// </summary>
        [Input("strategy")]
        public Input<string>? Strategy { get; set; }

        public SourceMicrosoftSharepointConfigurationStreamFormatDocumentFileTypeFormatExperimentalArgs()
        {
        }
        public static new SourceMicrosoftSharepointConfigurationStreamFormatDocumentFileTypeFormatExperimentalArgs Empty => new SourceMicrosoftSharepointConfigurationStreamFormatDocumentFileTypeFormatExperimentalArgs();
    }
}
