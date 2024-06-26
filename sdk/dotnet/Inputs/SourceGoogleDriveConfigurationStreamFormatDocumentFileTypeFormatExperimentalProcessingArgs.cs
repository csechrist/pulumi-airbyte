// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Airbyte.Inputs
{

    public sealed class SourceGoogleDriveConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Process files locally, supporting `fast` and `ocr` modes. This is the default option.
        /// </summary>
        [Input("local")]
        public Input<Inputs.SourceGoogleDriveConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingLocalArgs>? Local { get; set; }

        public SourceGoogleDriveConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingArgs()
        {
        }
        public static new SourceGoogleDriveConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingArgs Empty => new SourceGoogleDriveConfigurationStreamFormatDocumentFileTypeFormatExperimentalProcessingArgs();
    }
}
