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
    public sealed class DestinationPineconeConfigurationProcessingTextSplitterByMarkdownHeader
    {
        /// <summary>
        /// Level of markdown headers to split text fields by. Headings down to the specified level will be used as split points. Default: 1
        /// </summary>
        public readonly int? SplitLevel;

        [OutputConstructor]
        private DestinationPineconeConfigurationProcessingTextSplitterByMarkdownHeader(int? splitLevel)
        {
            SplitLevel = splitLevel;
        }
    }
}
