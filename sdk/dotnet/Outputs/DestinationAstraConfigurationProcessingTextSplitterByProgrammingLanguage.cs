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
    public sealed class DestinationAstraConfigurationProcessingTextSplitterByProgrammingLanguage
    {
        /// <summary>
        /// Split code in suitable places based on the programming language. must be one of ["cpp", "go", "java", "js", "php", "proto", "python", "rst", "ruby", "rust", "scala", "swift", "markdown", "latex", "html", "sol"]
        /// </summary>
        public readonly string Language;

        [OutputConstructor]
        private DestinationAstraConfigurationProcessingTextSplitterByProgrammingLanguage(string language)
        {
            Language = language;
        }
    }
}
